package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

var KKindKeys = []string{
	"bindings", "componentstatuses", "configmaps", "endpoints", "events",
	"limitranges", "namespaces", "nodes", "persistentvolumeclaims", "persistentvolumes",
	"pods", "podtemplates", "replicationcontrollers", "resourcequotas", "secrets",
	"secret", "serviceaccounts", "services", "challenges", "orders",
	"mutatingwebhookconfigurations", "validatingwebhookconfigurations", "customresourcedefinitions", "apiservices",
	"controllerrevisions", "daemonsets", "deployments", "replicasets", "statefulsets",
	"applications", "applicationsets", "appprojects", "argocdextensions", "tokenreviews",
	"localsubjectaccessreviews", "selfsubjectaccessreviews", "selfsubjectrulesreviews", "subjectaccessreviews",
	"horizontalpodautoscalers", "cronjobs", "jobs", "certificaterequests", "certificates",
	"clusterissuers", "issuers", "certificatesigningrequests", "devworkspaceoperatorconfigs", "devworkspaceroutings",
	"leases", "bgpconfigurations", "bgppeers", "blockaffinities", "caliconodestatuses",
	"clusterinformations", "felixconfigurations", "globalnetworkpolicies", "globalnetworksets", "hostendpoints",
	"ipamblocks", "ipamconfigs", "ipamhandles", "ippools", "ipreservations",
	"kubecontrollersconfigurations", "networkpolicies", "networksets", "endpointslices", "flowschemas",
	"prioritylevelconfigurations", "departments", "licenses", "products", "projects",
	"roletemplatebindings", "roletemplates", "users", "user", "workzooms",
	"workzoom", "clusterpolicies", "clusterreportchangerequests", "generaterequests", "policies",
	"reportchangerequests", "updaterequests", "alertmanagerconfigs", "alertmanagers", "podmonitors",
	"probes", "prometheuses", "prometheusrules", "servicemonitors", "thanosrulers",
	"ingressclasses", "ingresses", "ingress", "runtimeclasses", "checlusters",
	"poddisruptionbudgets", "clusterrolebindings", "clusterroles", "rolebindings", "roles",
	"priorityclasses", "csidrivers", "csinodes", "csistoragecapacities", "storageclasses",
	"volumeattachments", "clusterpolicyreports", "policyreports", "devworkspaces", "devworkspacetemplates",
	"cs", "cm", "ep", "ev", "limits", "ns", "no", "pvc", "pv", "po",
	"rc", "quota", "sa", "svc", "crd", "crds", "ds", "deploy", "rs", "sts",
	"app", "apps", "appset", "appsets", "appproj", "appprojs", "hpa", "cj", "cr", "crs",
	"cert", "certs", "csr", "dwoc", "dwr", "dep", "cpol", "crcr", "gr", "pol", "rcr", "ur",
	"amcfg", "am", "pmon", "prb", "prom", "promrule", "smon", "ruler", "ing", "netpol", "pdb", "pc", "sc",
	"cpolr", "polr", "dw", "dwt", "-n",
}

var KArgValues map[string]*c.TreeNode

func KKind() map[string]*c.TreeNode {
	if KArgValues == nil {
		KArgValues = make(map[string]*c.TreeNode, len(KKindKeys))
		for _, v := range KKindKeys {
			KArgValues[v] = &c.TreeNode{
				IsShell:  true,
				Shell:    "kubectl get " + v + " -A",
				Children: KArgs(0),
			}
		}
	}
	return KArgValues
}

var KArgsValues = map[string]string{
	"-n":  "ns",
	"-A":  "ns -A",
	"-f":  "po -A",
	"-c":  `-o=jsonpath='{range .spec.containers[*]}{.name}{"\n"}{end}'`,
	"-it": "po -A",
	"-o":  "yaml",
}

var KAArgsValues map[string]*c.TreeNode

func KArgs(level int) map[string]*c.TreeNode {
	if level == 4 {
		return nil
	}
	level++
	if KAArgsValues == nil {
		KAArgsValues = make(map[string]*c.TreeNode)
		for k, v := range KArgsValues {
			if k == "-c" {
				KAArgsValues[k] = &c.TreeNode{
					IsShell:  true,
					Shell:    v,
					Children: KArgs(level),
				}
				continue
			}
			KAArgsValues[k] = &c.TreeNode{
				IsShell:  true,
				Shell:    "kubectl get " + v,
				Children: KArgs(level),
			}
		}
	}
	return KAArgsValues
}
