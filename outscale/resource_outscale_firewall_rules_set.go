package outscale

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
)

func resourceOutscaleFirewallRulesSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutscaleSecurityGroupCreate,
		Read:   resourceOutscaleSecurityGroupRead,
		Update: resourceOutscaleSecurityGroupUpdate,
		Delete: resourceOutscaleSecurityGroupDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name_prefix"},
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) > 255 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 255 characters", k))
					}
					return
				},
			},

			"name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) > 100 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 100 characters, group_name is limited to 255", k))
					}
					return
				},
			},

			"group_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "Managed by Terraform",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) > 255 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 255 characters", k))
					}
					return
				},
			},

			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"ip_permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from_port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"to_port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"ip_protocol": {
							Type:      schema.TypeString,
							Required:  true,
							StateFunc: protocolStateFunc,
						},

						"ip_ranges": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validateCIDRNetworkAddress,
							},
						},

						"groups": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},

						"self": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
				Set: resourceOutscaleSecurityGroupRuleHash,
			},

			"ip_permissions_egress": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from_port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"to_port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"ip_protocol": {
							Type:      schema.TypeString,
							Required:  true,
							StateFunc: protocolStateFunc,
						},

						"ip_ranges": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validateCIDRNetworkAddress,
							},
						},

						"prefix_list_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"groups": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},

						"self": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
				Set: resourceOutscaleSecurityGroupRuleHash,
			},

			"owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tags":    tagsSchema(),
			"tag_set": tagsSchemaComputed(),
		},
	}
}

func resourceOutscaleSecurityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	securityGroupOpts := &fcu.CreateSecurityGroupInput{}

	if v, ok := d.GetOk("vpc_id"); ok {
		securityGroupOpts.VpcId = aws.String(v.(string))
	}

	if v := d.Get("group_description"); v != nil {
		securityGroupOpts.Description = aws.String(v.(string))
	}

	var groupName string
	if v, ok := d.GetOk("group_name"); ok {
		groupName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		groupName = resource.PrefixedUniqueId(v.(string))
	} else {
		groupName = resource.UniqueId()
	}
	securityGroupOpts.GroupName = aws.String(groupName)

	var err error
	log.Printf(
		"[DEBUG] Security Group create configuration: %#v", securityGroupOpts)

	var createResp *fcu.CreateSecurityGroupOutput

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		createResp, err = conn.VM.CreateSecurityGroup(securityGroupOpts)

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("Error creating Security Group: %s", err)
	}

	d.SetId(*createResp.GroupId)

	log.Printf("[INFO] Security Group ID: %s", d.Id())

	// Wait for the security group to truly exist
	log.Printf(
		"[DEBUG] Waiting for Security Group (%s) to exist",
		d.Id())
	stateConf := &resource.StateChangeConf{
		Pending: []string{""},
		Target:  []string{"exists"},
		Refresh: SGStateRefreshFunc(conn, d.Id()),
		Timeout: 3 * time.Minute,
	}

	resp, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for Security Group (%s) to become available: %s",
			d.Id(), err)
	}

	if err := setTags(conn, d); err != nil {
		return err
	}

	// AWS defaults all Security Groups to have an ALLOW ALL egress rule. Here we
	// revoke that rule, so users don't unknowingly have/use it.
	group := resp.(*fcu.SecurityGroup)
	if group.VpcId != nil && *group.VpcId != "" {
		log.Printf("[DEBUG] Revoking default egress rule for Security Group for %s", d.Id())

		req := &fcu.RevokeSecurityGroupEgressInput{
			GroupId: createResp.GroupId,
			IpPermissions: []*fcu.IpPermission{
				{
					FromPort: aws.Int64(int64(0)),
					ToPort:   aws.Int64(int64(0)),
					IpRanges: []*fcu.IpRange{
						{
							CidrIp: aws.String("0.0.0.0/0"),
						},
					},
					IpProtocol: aws.String("-1"),
				},
			},
		}

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			_, err = conn.VM.RevokeSecurityGroupEgress(req)

			if err != nil {
				if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf(
				"Error revoking default egress rule for Security Group (%s): %s",
				d.Id(), err)
		}

	}

	return resourceOutscaleSecurityGroupUpdate(d, meta)
}

func resourceOutscaleSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	sgRaw, _, err := SGStateRefreshFunc(conn, d.Id())()
	if err != nil {
		return err
	}
	if sgRaw == nil {
		d.SetId("")
		return nil
	}

	sg := sgRaw.(*fcu.SecurityGroup)

	remoteIngressRules := resourceOutscaleSecurityGroupIPPermGather(d.Id(), sg.IpPermissions, sg.OwnerId)
	remoteEgressRules := resourceOutscaleSecurityGroupIPPermGather(d.Id(), sg.IpPermissionsEgress, sg.OwnerId)

	// localIngressRules := d.Get("ip_permissions").(*schema.Set).List()
	// localEgressRules := d.Get("ip_permissions_egress").(*schema.Set).List()

	// Loop through the local state of rules, doing a match against the remote
	// ruleSet we built above.
	// ingressRules := matchRules("ip_permissions", localIngressRules, remoteIngressRules)
	// egressRules := matchRules("ip_permissions_egress", localEgressRules, remoteEgressRules)

	d.Set("group_description", sg.Description)
	d.Set("group_name", sg.GroupName)
	d.Set("vpc_id", sg.VpcId)
	d.Set("owner_id", sg.OwnerId)

	if err := d.Set("ip_permissions", remoteIngressRules); err != nil {
		log.Printf("[WARN] Error setting Ingress rule set for (%s): %s", d.Id(), err)
	}

	if err := d.Set("ip_permissions_egress", remoteEgressRules); err != nil {
		log.Printf("[WARN] Error setting Egress rule set for (%s): %s", d.Id(), err)
	}

	d.Set("tag_set", tagsToMap(sg.Tags))
	return nil
}

func resourceOutscaleSecurityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	sgRaw, _, err := SGStateRefreshFunc(conn, d.Id())()
	if err != nil {
		return err
	}
	if sgRaw == nil {
		d.SetId("")
		return nil
	}

	group := sgRaw.(*fcu.SecurityGroup)

	err = resourceOutscaleSecurityGroupUpdateRules(d, "ip_permissions", meta, group)
	if err != nil {
		return err
	}

	if d.Get("vpc_id") != nil {
		err = resourceOutscaleSecurityGroupUpdateRules(d, "ip_permissions_egress", meta, group)
		if err != nil {
			return err
		}
	}

	if !d.IsNewResource() {
		if err := setTags(conn, d); err != nil {
			return err
		}
		d.SetPartial("tag_set")
	}

	return resourceOutscaleSecurityGroupRead(d, meta)
}

func resourceOutscaleSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	log.Printf("[DEBUG] Security Group destroy: %v", d.Id())

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		_, err := conn.VM.DeleteSecurityGroup(&fcu.DeleteSecurityGroupInput{
			GroupId: aws.String(d.Id()),
		})
		if err != nil {
			ec2err, ok := err.(awserr.Error)
			if !ok {
				return resource.RetryableError(err)
			}

			switch ec2err.Code() {
			case "InvalidGroup.NotFound":
				return nil
			case "DependencyViolation":
				// If it is a dependency violation, we want to retry
				return resource.RetryableError(err)
			default:
				// Any other error, we want to quit the retry loop immediately
				return resource.NonRetryableError(err)
			}
		}

		return nil
	})
}

func resourceOutscaleSecurityGroupRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%d-", m["from_port"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["to_port"].(int)))
	p := protocolForValue(m["ip_protocol"].(string))
	buf.WriteString(fmt.Sprintf("%s-", p))
	buf.WriteString(fmt.Sprintf("%t-", m["self"].(bool)))

	// We need to make sure to sort the strings below so that we always
	// generate the same hash code no matter what is in the set.
	if v, ok := m["ip_ranges"]; ok {
		vs := v.([]interface{})
		s := make([]string, len(vs))
		for i, raw := range vs {
			s[i] = raw.(string)
		}
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s-", v))
		}
	}
	if v, ok := m["prefix_list_ids"]; ok {
		vs := v.([]interface{})
		s := make([]string, len(vs))
		for i, raw := range vs {
			s[i] = raw.(string)
		}
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s-", v))
		}
	}
	if v, ok := m["groups"]; ok {
		vs := v.(*schema.Set).List()
		s := make([]string, len(vs))
		for i, raw := range vs {
			s[i] = raw.(string)
		}
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s-", v))
		}
	}

	return hashcode.String(buf.String())
}

func resourceOutscaleSecurityGroupIPPermGather(groupId string, permissions []*fcu.IpPermission, ownerId *string) []map[string]interface{} {
	ruleMap := make(map[string]map[string]interface{})
	for _, perm := range permissions {
		var fromPort, toPort int64
		if v := perm.FromPort; v != nil {
			fromPort = *v
		}
		if v := perm.ToPort; v != nil {
			toPort = *v
		}

		k := fmt.Sprintf("%s-%d-%d", *perm.IpProtocol, fromPort, toPort)
		m, ok := ruleMap[k]
		if !ok {
			m = make(map[string]interface{})
			ruleMap[k] = m
		}

		m["from_port"] = fromPort
		m["to_port"] = toPort
		m["ip_protocol"] = *perm.IpProtocol

		if len(perm.IpRanges) > 0 {
			raw, ok := m["ip_ranges"]
			if !ok {
				raw = make([]string, 0, len(perm.IpRanges))
			}
			list := raw.([]string)

			for _, ip := range perm.IpRanges {
				list = append(list, *ip.CidrIp)
			}

			m["ip_ranges"] = list
		}

		if len(perm.PrefixListIds) > 0 {
			raw, ok := m["prefix_list_ids"]
			if !ok {
				raw = make([]string, 0, len(perm.PrefixListIds))
			}
			list := raw.([]string)

			for _, pl := range perm.PrefixListIds {
				list = append(list, *pl.PrefixListId)
			}

			m["prefix_list_ids"] = list
		}

		groups := flattenSecurityGroups(perm.UserIdGroupPairs, ownerId)
		for i, g := range groups {
			if *g.GroupId == groupId {
				groups[i], groups = groups[len(groups)-1], groups[:len(groups)-1]
				m["self"] = true
			}
		}

		if len(groups) > 0 {
			raw, ok := m["groups"]
			if !ok {
				raw = schema.NewSet(schema.HashString, nil)
			}
			list := raw.(*schema.Set)

			for _, g := range groups {
				if g.GroupName != nil {
					list.Add(*g.GroupName)
				} else {
					list.Add(*g.GroupId)
				}
			}

			m["groups"] = list
		}
	}
	rules := make([]map[string]interface{}, 0, len(ruleMap))
	for _, m := range ruleMap {
		rules = append(rules, m)
	}

	return rules
}

func resourceOutscaleSecurityGroupUpdateRules(
	d *schema.ResourceData, ruleset string,
	meta interface{}, group *fcu.SecurityGroup) error {

	if d.HasChange(ruleset) {
		o, n := d.GetChange(ruleset)
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		remove, err := expandIPPerms(group, os.Difference(ns).List())
		if err != nil {
			return err
		}
		add, err := expandIPPerms(group, ns.Difference(os).List())
		if err != nil {
			return err
		}

		if len(remove) > 0 || len(add) > 0 {
			conn := meta.(*OutscaleClient).FCU

			var err error
			if len(remove) > 0 {
				log.Printf("[DEBUG] Revoking security group %#v %s rule: %#v",
					group, ruleset, remove)

				if ruleset == "egress" {
					req := &fcu.RevokeSecurityGroupEgressInput{
						GroupId:       group.GroupId,
						IpPermissions: remove,
					}

					err = resource.Retry(5*time.Minute, func() *resource.RetryError {
						_, err = conn.VM.RevokeSecurityGroupEgress(req)

						if err != nil {
							if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
								return resource.RetryableError(err)
							}
							return resource.NonRetryableError(err)
						}

						return nil
					})

				} else {
					req := &fcu.RevokeSecurityGroupIngressInput{
						GroupId:       group.GroupId,
						IpPermissions: remove,
					}
					if group.VpcId == nil || *group.VpcId == "" {
						req.GroupId = nil
						req.GroupName = group.GroupName
					}
					err = resource.Retry(5*time.Minute, func() *resource.RetryError {
						_, err = conn.VM.RevokeSecurityGroupIngress(req)

						if err != nil {
							if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
								return resource.RetryableError(err)
							}
							return resource.NonRetryableError(err)
						}

						return nil
					})
				}

				if err != nil {
					return fmt.Errorf(
						"Error revoking security group %s rules: %s",
						ruleset, err)
				}
			}

			if len(add) > 0 {
				log.Printf("[DEBUG] Authorizing security group %#v %s rule: %#v",
					group, ruleset, add)
				// Authorize the new rules
				if ruleset == "egress" {
					req := &fcu.AuthorizeSecurityGroupEgressInput{
						GroupId:       group.GroupId,
						IpPermissions: add,
					}

					err = resource.Retry(5*time.Minute, func() *resource.RetryError {
						_, err = conn.VM.AuthorizeSecurityGroupEgress(req)

						if err != nil {
							if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
								return resource.RetryableError(err)
							}
							return resource.NonRetryableError(err)
						}

						return nil
					})
				} else {
					req := &fcu.AuthorizeSecurityGroupIngressInput{
						GroupId:       group.GroupId,
						IpPermissions: add,
					}
					if group.VpcId == nil || *group.VpcId == "" {
						req.GroupId = nil
						req.GroupName = group.GroupName
					}

					err = resource.Retry(5*time.Minute, func() *resource.RetryError {
						_, err = conn.VM.AuthorizeSecurityGroupIngress(req)

						if err != nil {
							if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
								return resource.RetryableError(err)
							}
							return resource.NonRetryableError(err)
						}

						return nil
					})
				}

				if err != nil {
					return fmt.Errorf(
						"Error authorizing security group %s rules: %s",
						ruleset, err)
				}
			}
		}
	}
	return nil
}

// SGStateRefreshFunc returns a resource.StateRefreshFunc that is used to watch
// a security group.
func SGStateRefreshFunc(conn *fcu.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		req := &fcu.DescribeSecurityGroupsInput{
			GroupIds: []*string{aws.String(id)},
		}

		var resp *fcu.DescribeSecurityGroupsOutput
		var err error

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			resp, err = conn.VM.DescribeSecurityGroups(req)

			if err != nil {
				if strings.Contains(err.Error(), "RequestLimitExceeded") || strings.Contains(err.Error(), "DependencyViolation") {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if err != nil {
			if ec2err, ok := err.(awserr.Error); ok {
				if ec2err.Code() == "InvalidSecurityGroupID.NotFound" ||
					ec2err.Code() == "InvalidGroup.NotFound" {
					resp = nil
					err = nil
				}
			}

			if err != nil {
				log.Printf("Error on SGStateRefresh: %s", err)
				return nil, "", err
			}
		}

		if resp == nil {
			return nil, "", nil
		}

		group := resp.SecurityGroups[0]
		return group, "exists", nil
	}
}

func matchRules(rType string, local []interface{}, remote []map[string]interface{}) []map[string]interface{} {

	var saves []map[string]interface{}
	for _, raw := range local {
		l := raw.(map[string]interface{})

		var selfVal bool
		if v, ok := l["self"]; ok {
			selfVal = v.(bool)
		}

		localHash := idHash(rType, l["ip_protocol"].(string), int64(l["to_port"].(int)), int64(l["from_port"].(int)), selfVal)

		for _, r := range remote {
			var remoteSelfVal bool
			if v, ok := r["self"]; ok {
				remoteSelfVal = v.(bool)
			}

			rHash := idHash(rType, r["ip_protocol"].(string), r["to_port"].(int64), r["from_port"].(int64), remoteSelfVal)
			if rHash == localHash {
				var numExpectedCidrs, numExpectedPrefixLists, numExpectedSGs, numRemoteCidrs, numRemotePrefixLists, numRemoteSGs int
				var matchingCidrs []string
				var matchingSGs []string
				var matchingPrefixLists []string

				lcRaw, ok := l["ip_ranges"]
				if ok {
					numExpectedCidrs = len(l["ip_ranges"].([]interface{}))
				}

				lpRaw, ok := l["prefix_list_ids"]
				if ok {
					numExpectedPrefixLists = len(l["prefix_list_ids"].([]interface{}))
				}
				lsRaw, ok := l["groups"]
				if ok {
					numExpectedSGs = len(l["groups"].(*schema.Set).List())
				}

				rcRaw, ok := r["ip_ranges"]
				if ok {
					numRemoteCidrs = len(r["ip_ranges"].([]string))
				}
				rpRaw, ok := r["prefix_list_ids"]
				if ok {
					numRemotePrefixLists = len(r["prefix_list_ids"].([]string))
				}

				rsRaw, ok := r["groups"]
				if ok {
					numRemoteSGs = len(r["groups"].(*schema.Set).List())
				}

				// check some early failures
				if numExpectedCidrs > numRemoteCidrs {
					log.Printf("[DEBUG] Local rule has more CIDR blocks, continuing (%d/%d)", numExpectedCidrs, numRemoteCidrs)
					continue
				}

				if numExpectedPrefixLists > numRemotePrefixLists {
					log.Printf("[DEBUG] Local rule has more prefix lists, continuing (%d/%d)", numExpectedPrefixLists, numRemotePrefixLists)
					continue
				}
				if numExpectedSGs > numRemoteSGs {
					log.Printf("[DEBUG] Local rule has more Security Groups, continuing (%d/%d)", numExpectedSGs, numRemoteSGs)
					continue
				}

				// match CIDRs by converting both to sets, and using Set methods
				var localCidrs []interface{}
				if lcRaw != nil {
					localCidrs = lcRaw.([]interface{})
				}
				localCidrSet := schema.NewSet(schema.HashString, localCidrs)

				// remote cidrs are presented as a slice of strings, so we need to
				// reformat them into a slice of interfaces to be used in creating the
				// remote cidr set
				var remoteCidrs []string
				if rcRaw != nil {
					remoteCidrs = rcRaw.([]string)
				}
				// convert remote cidrs to a set, for easy comparisons
				var list []interface{}
				for _, s := range remoteCidrs {
					list = append(list, s)
				}
				remoteCidrSet := schema.NewSet(schema.HashString, list)

				// Build up a list of local cidrs that are found in the remote set
				for _, s := range localCidrSet.List() {
					if remoteCidrSet.Contains(s) {
						matchingCidrs = append(matchingCidrs, s.(string))
					}
				}

				var localPrefixLists []interface{}
				if lpRaw != nil {
					localPrefixLists = lpRaw.([]interface{})
				}
				localPrefixListsSet := schema.NewSet(schema.HashString, localPrefixLists)

				// remote prefix lists are presented as a slice of strings, so we need to
				// reformat them into a slice of interfaces to be used in creating the
				// remote prefix list set
				var remotePrefixLists []string
				if rpRaw != nil {
					remotePrefixLists = rpRaw.([]string)
				}
				// convert remote prefix lists to a set, for easy comparison
				list = nil
				for _, s := range remotePrefixLists {
					list = append(list, s)
				}
				remotePrefixListsSet := schema.NewSet(schema.HashString, list)

				// Build up a list of local prefix lists that are found in the remote set
				for _, s := range localPrefixListsSet.List() {
					if remotePrefixListsSet.Contains(s) {
						matchingPrefixLists = append(matchingPrefixLists, s.(string))
					}
				}

				// match SGs. Both local and remote are already sets
				var localSGSet *schema.Set
				if lsRaw == nil {
					localSGSet = schema.NewSet(schema.HashString, nil)
				} else {
					localSGSet = lsRaw.(*schema.Set)
				}

				var remoteSGSet *schema.Set
				if rsRaw == nil {
					remoteSGSet = schema.NewSet(schema.HashString, nil)
				} else {
					remoteSGSet = rsRaw.(*schema.Set)
				}

				// Build up a list of local security groups that are found in the remote set
				for _, s := range localSGSet.List() {
					if remoteSGSet.Contains(s) {
						matchingSGs = append(matchingSGs, s.(string))
					}
				}

				// compare equalities for matches.
				// If we found the number of cidrs and number of sgs, we declare a
				// match, and then remove those elements from the remote rule, so that
				// this remote rule can still be considered by other local rules
				if numExpectedCidrs == len(matchingCidrs) {
					if numExpectedPrefixLists == len(matchingPrefixLists) {
						if numExpectedSGs == len(matchingSGs) {
							// confirm that self references match
							var lSelf bool
							var rSelf bool
							if _, ok := l["self"]; ok {
								lSelf = l["self"].(bool)
							}
							if _, ok := r["self"]; ok {
								rSelf = r["self"].(bool)
							}
							if rSelf == lSelf {
								delete(r, "self")
								// pop local cidrs from remote
								diffCidr := remoteCidrSet.Difference(localCidrSet)
								var newCidr []string
								for _, cRaw := range diffCidr.List() {
									newCidr = append(newCidr, cRaw.(string))
								}

								// reassigning
								if len(newCidr) > 0 {
									r["ip_ranges"] = newCidr
								} else {
									delete(r, "ip_ranges")
								}

								// pop local prefix lists from remote
								diffPrefixLists := remotePrefixListsSet.Difference(localPrefixListsSet)
								var newPrefixLists []string
								for _, pRaw := range diffPrefixLists.List() {
									newPrefixLists = append(newPrefixLists, pRaw.(string))
								}

								// reassigning
								if len(newPrefixLists) > 0 {
									r["prefix_list_ids"] = newPrefixLists
								} else {
									delete(r, "prefix_list_ids")
								}

								// pop local sgs from remote
								diffSGs := remoteSGSet.Difference(localSGSet)
								if len(diffSGs.List()) > 0 {
									r["groups"] = diffSGs
								} else {
									delete(r, "groups")
								}

								saves = append(saves, l)
							}
						}
					}

				}
			}
		}
	}
	// Here we catch any remote rules that have not been stripped of all self,
	// cidrs, and security groups. We'll add remote rules here that have not been
	// matched locally, and let the graph sort things out. This will happen when
	// rules are added externally to Terraform
	for _, r := range remote {
		var lenCidr, lenPrefixLists, lenSGs int
		if rCidrs, ok := r["ip_ranges"]; ok {
			lenCidr = len(rCidrs.([]string))
		}
		if rPrefixLists, ok := r["prefix_list_ids"]; ok {
			lenPrefixLists = len(rPrefixLists.([]string))
		}
		if rawSGs, ok := r["groups"]; ok {
			lenSGs = len(rawSGs.(*schema.Set).List())
		}

		if _, ok := r["self"]; ok {
			if r["self"].(bool) == true {
				lenSGs++
			}
		}

		if lenSGs+lenCidr+lenPrefixLists > 0 {
			log.Printf("[DEBUG] Found a remote Rule that wasn't empty: (%#v)", r)
			saves = append(saves, r)
		}
	}

	return saves
}

// Creates a unique hash for the type, ports, and ip_protocol, used as a key in
// maps
func idHash(rType, ip_protocol string, toPort, fromPort int64, self bool) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s-", rType))
	buf.WriteString(fmt.Sprintf("%d-", toPort))
	buf.WriteString(fmt.Sprintf("%d-", fromPort))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(ip_protocol)))
	buf.WriteString(fmt.Sprintf("%t-", self))

	return fmt.Sprintf("rule-%d", hashcode.String(buf.String()))
}

// protocolStateFunc ensures we only store a string in any ip_protocol field
func protocolStateFunc(v interface{}) string {
	switch v.(type) {
	case string:
		p := protocolForValue(v.(string))
		return p
	default:
		log.Printf("[WARN] Non String value given for Protocol: %#v", v)
		return ""
	}
}

// protocolForValue converts a valid Internet Protocol number into it's group_name
// representation. If a group_name is given, it validates that it's a proper ip_protocol
// group_name. Names/numbers are as defined at
// https://www.iana.org/assignments/ip_protocol-numbers/ip_protocol-numbers.xhtml
func protocolForValue(v string) string {
	// special case -1
	ip_protocol := strings.ToLower(v)
	if ip_protocol == "-1" || ip_protocol == "all" {
		return "-1"
	}
	// if it's a group_name like tcp, return that
	if _, ok := sgProtocolIntegers()[ip_protocol]; ok {
		return ip_protocol
	}
	// convert to int, look for that value
	p, err := strconv.Atoi(ip_protocol)
	if err != nil {
		// we were unable to convert to int, suggesting a string group_name, but it wasn't
		// found above
		log.Printf("[WARN] Unable to determine valid ip_protocol: %s", err)
		return ip_protocol
	}

	for k, v := range sgProtocolIntegers() {
		if p == v {
			// guard against protocolIntegers sometime in the future not having lower
			// case ids in the map
			return strings.ToLower(k)
		}
	}

	// fall through
	log.Printf("[WARN] Unable to determine valid ip_protocol: no matching protocols found")
	return ip_protocol
}

// a map of ip_protocol names and their codes, defined at
// https://www.iana.org/assignments/ip_protocol-numbers/ip_protocol-numbers.xhtml,
// documented to be supported by AWS Security Groups
// http://docs.aws.amazon.com/fr_fr/AWSEC2/latest/APIReference/API_IpPermission.html
// Similar to protocolIntegers() used by Network ACLs, but explicitly only
// supports "tcp", "udp", "icmp", and "all"
func sgProtocolIntegers() map[string]int {
	var protocolIntegers = make(map[string]int)
	protocolIntegers = map[string]int{
		"udp":  17,
		"tcp":  6,
		"icmp": 1,
		"all":  -1,
	}
	return protocolIntegers
}

func flattenSecurityGroups(list []*fcu.UserIdGroupPair, ownerId *string) []*fcu.GroupIdentifier {
	result := make([]*fcu.GroupIdentifier, 0, len(list))
	for _, g := range list {
		var userId *string
		if g.UserId != nil && *g.UserId != "" && (ownerId == nil || *ownerId != *g.UserId) {
			userId = g.UserId
		}
		// userid nil here for same vpc groups

		vpc := g.GroupName == nil || *g.GroupName == ""
		var id *string
		if vpc {
			id = g.GroupId
		} else {
			id = g.GroupName
		}

		// id is groupid for vpcs
		// id is groupname for non vpc (classic)

		if userId != nil {
			id = aws.String(*userId + "/" + *id)
		}

		if vpc {
			result = append(result, &fcu.GroupIdentifier{
				GroupId: id,
			})
		} else {
			result = append(result, &fcu.GroupIdentifier{
				GroupId:   g.GroupId,
				GroupName: id,
			})
		}
	}
	return result
}

func expandIPPerms(
	group *fcu.SecurityGroup, configured []interface{}) ([]*fcu.IpPermission, error) {
	vpc := group.VpcId != nil && *group.VpcId != ""

	perms := make([]*fcu.IpPermission, len(configured))
	for i, mRaw := range configured {
		var perm fcu.IpPermission
		m := mRaw.(map[string]interface{})

		perm.FromPort = aws.Int64(int64(m["from_port"].(int)))
		perm.ToPort = aws.Int64(int64(m["to_port"].(int)))
		perm.IpProtocol = aws.String(m["ip_protocol"].(string))

		// When ip_protocol is "-1", AWS won't store any ports for the
		// rule, but also won't error if the user specifies ports other
		// than '0'. Force the user to make a deliberate '0' port
		// choice when specifying a "-1" ip_protocol, and tell them about
		// AWS's behavior in the error message.
		if *perm.IpProtocol == "-1" && (*perm.FromPort != 0 || *perm.ToPort != 0) {
			return nil, fmt.Errorf(
				"from_port (%d) and to_port (%d) must both be 0 to use the 'ALL' \"-1\" ip_protocol!",
				*perm.FromPort, *perm.ToPort)
		}

		var groups []string
		if raw, ok := m["groups"]; ok {
			list := raw.(*schema.Set).List()
			for _, v := range list {
				groups = append(groups, v.(string))
			}
		}
		if v, ok := m["self"]; ok && v.(bool) {
			if vpc {
				groups = append(groups, *group.GroupId)
			} else {
				groups = append(groups, *group.GroupName)
			}
		}

		if len(groups) > 0 {
			perm.UserIdGroupPairs = make([]*fcu.UserIdGroupPair, len(groups))
			for i, name := range groups {
				ownerId, id := "", name
				if items := strings.Split(id, "/"); len(items) > 1 {
					ownerId, id = items[0], items[1]
				}

				perm.UserIdGroupPairs[i] = &fcu.UserIdGroupPair{
					GroupId: aws.String(id),
				}

				if ownerId != "" {
					perm.UserIdGroupPairs[i].UserId = aws.String(ownerId)
				}

				if !vpc {
					perm.UserIdGroupPairs[i].GroupId = nil
					perm.UserIdGroupPairs[i].GroupName = aws.String(id)
				}
			}
		}

		if raw, ok := m["ip_ranges"]; ok {
			list := raw.([]interface{})
			for _, v := range list {
				perm.IpRanges = append(perm.IpRanges, &fcu.IpRange{CidrIp: aws.String(v.(string))})
			}
		}
		if raw, ok := m["prefix_list_ids"]; ok {
			list := raw.([]interface{})
			for _, v := range list {
				perm.PrefixListIds = append(perm.PrefixListIds, &fcu.PrefixListId{PrefixListId: aws.String(v.(string))})
			}
		}

		perms[i] = &perm
	}

	return perms, nil
}
