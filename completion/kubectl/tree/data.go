package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

// k get po 第三极数据
// k get -n
var KKindValues = map[string]string{
	"bindings":                        "bindings",
	"componentstatuses":               "componentstatuses",
	"configmaps":                      "configmaps",
	"endpoints":                       "endpoints",
	"events":                          "events",
	"limitranges":                     "limitranges",
	"namespaces":                      "namespaces",
	"nodes":                           "nodes",
	"persistentvolumeclaims":          "persistentvolumeclaims",
	"persistentvolumes":               "persistentvolumes",
	"pods":                            "pods",
	"podtemplates":                    "podtemplates",
	"replicationcontrollers":          "replicationcontrollers",
	"resourcequotas":                  "resourcequotas",
	"secrets":                         "secrets",
	"secret":                          "secret",
	"serviceaccounts":                 "serviceaccounts",
	"services":                        "services",
	"challenges":                      "challenges",
	"orders":                          "orders",
	"mutatingwebhookconfigurations":   "mutatingwebhookconfigurations",
	"validatingwebhookconfigurations": "validatingwebhookconfigurations",
	"customresourcedefinitions":       "customresourcedefinitions",
	"apiservices":                     "apiservices",
	"controllerrevisions":             "controllerrevisions",
	"daemonsets":                      "daemonsets",
	"deployments":                     "deployments",
	"replicasets":                     "replicasets",
	"statefulsets":                    "statefulsets",
	"applications":                    "applications",
	"applicationsets":                 "applicationsets",
	"appprojects":                     "appprojects",
	"argocdextensions":                "argocdextensions",
	"tokenreviews":                    "tokenreviews",
	"localsubjectaccessreviews":       "localsubjectaccessreviews",
	"selfsubjectaccessreviews":        "selfsubjectaccessreviews",
	"selfsubjectrulesreviews":         "selfsubjectrulesreviews",
	"subjectaccessreviews":            "subjectaccessreviews",
	"horizontalpodautoscalers":        "horizontalpodautoscalers",
	"cronjobs":                        "cronjobs",
	"jobs":                            "jobs",
	"certificaterequests":             "certificaterequests",
	"certificates":                    "certificates",
	"clusterissuers":                  "clusterissuers",
	"issuers":                         "issuers",
	"certificatesigningrequests":      "certificatesigningrequests",
	"devworkspaceoperatorconfigs":     "devworkspaceoperatorconfigs",
	"devworkspaceroutings":            "devworkspaceroutings",
	"leases":                          "leases",
	"bgpconfigurations":               "bgpconfigurations",
	"bgppeers":                        "bgppeers",
	"blockaffinities":                 "blockaffinities",
	"caliconodestatuses":              "caliconodestatuses",
	"clusterinformations":             "clusterinformations",
	"felixconfigurations":             "felixconfigurations",
	"globalnetworkpolicies":           "globalnetworkpolicies",
	"globalnetworksets":               "globalnetworksets",
	"hostendpoints":                   "hostendpoints",
	"ipamblocks":                      "ipamblocks",
	"ipamconfigs":                     "ipamconfigs",
	"ipamhandles":                     "ipamhandles",
	"ippools":                         "ippools",
	"ipreservations":                  "ipreservations",
	"kubecontrollersconfigurations":   "kubecontrollersconfigurations",
	"networkpolicies":                 "networkpolicies",
	"networksets":                     "networksets",
	"endpointslices":                  "endpointslices",
	"flowschemas":                     "flowschemas",
	"prioritylevelconfigurations":     "prioritylevelconfigurations",
	"departments":                     "departments",
	"licenses":                        "licenses",
	"products":                        "products",
	"projects":                        "projects",
	"roletemplatebindings":            "roletemplatebindings",
	"roletemplates":                   "roletemplates",
	"users":                           "users",
	"user":                            "user",
	"workzooms":                       "workzooms",
	"workzoom":                        "workzoom",
	"clusterpolicies":                 "clusterpolicies",
	"clusterreportchangerequests":     "clusterreportchangerequests",
	"generaterequests":                "generaterequests",
	"policies":                        "policies",
	"reportchangerequests":            "reportchangerequests",
	"updaterequests":                  "updaterequests",
	"alertmanagerconfigs":             "alertmanagerconfigs",
	"alertmanagers":                   "alertmanagers",
	"podmonitors":                     "podmonitors",
	"probes":                          "probes",
	"prometheuses":                    "prometheuses",
	"prometheusrules":                 "prometheusrules",
	"servicemonitors":                 "servicemonitors",
	"thanosrulers":                    "thanosrulers",
	"ingressclasses":                  "ingressclasses",
	"ingresses":                       "ingresses",
	"ingress":                         "ingress",
	"runtimeclasses":                  "runtimeclasses",
	"checlusters":                     "checlusters",
	"poddisruptionbudgets":            "poddisruptionbudgets",
	"clusterrolebindings":             "clusterrolebindings",
	"clusterroles":                    "clusterroles",
	"rolebindings":                    "rolebindings",
	"roles":                           "roles",
	"priorityclasses":                 "priorityclasses",
	"csidrivers":                      "csidrivers",
	"csinodes":                        "csinodes",
	"csistoragecapacities":            "csistoragecapacities",
	"storageclasses":                  "storageclasses",
	"volumeattachments":               "volumeattachments",
	"clusterpolicyreports":            "clusterpolicyreports",
	"policyreports":                   "policyreports",
	"devworkspaces":                   "devworkspaces",
	"devworkspacetemplates":           "devworkspacetemplates",
	"cs":                              "cs",
	"cm":                              "cm",
	"ep":                              "ep",
	"ev":                              "ev",
	"limits":                          "limits",
	"ns":                              "ns",
	"no":                              "no",
	"pvc":                             "pvc",
	"pv":                              "pv",
	"po":                              "po",
	"rc":                              "rc",
	"quota":                           "quota",
	"sa":                              "sa",
	"svc":                             "svc",
	"crd":                             "crd",
	"crds":                            "crds",
	"ds":                              "ds",
	"deploy":                          "deploy",
	"rs":                              "rs",
	"sts":                             "sts",
	"app":                             "app",
	"apps":                            "apps",
	"appset":                          "appset",
	"appsets":                         "appsets",
	"appproj":                         "appproj",
	"appprojs":                        "appprojs",
	"hpa":                             "hpa",
	"cj":                              "cj",
	"cr":                              "cr",
	"crs":                             "crs",
	"cert":                            "cert",
	"certs":                           "certs",
	"csr":                             "csr",
	"dwoc":                            "dwoc",
	"dwr":                             "dwr",
	"dep":                             "dep",
	"cpol":                            "cpol",
	"crcr":                            "crcr",
	"gr":                              "gr",
	"pol":                             "pol",
	"rcr":                             "rcr",
	"ur":                              "ur",
	"amcfg":                           "amcfg",
	"am":                              "am",
	"pmon":                            "pmon",
	"prb":                             "prb",
	"prom":                            "prom",
	"promrule":                        "promrule",
	"smon":                            "smon",
	"ruler":                           "ruler",
	"ing":                             "ing",
	"netpol":                          "netpol",
	"pdb":                             "pdb",
	"pc":                              "pc",
	"sc":                              "sc",
	"cpolr":                           "cpolr",
	"polr":                            "polr",
	"dw":                              "dw",
	"dwt":                             "dwt",
	"-n":                              "ns",
}

var KArgValues map[string]*c.TreeNode

func KKind() map[string]*c.TreeNode {
	if KArgValues == nil {
		KArgValues = make(map[string]*c.TreeNode)
		for k, v := range KKindValues {
			KArgValues[k] = &c.TreeNode{
				IsShell:  true,
				Shell:    "kubectl get " + v + " -A",
				Children: KArgs(0),
			}
		}
	}
	return KArgValues
}

// k get -n
// k get po -n
var KArgsValues = map[string]string{
	"-n":  "ns",
	"-A":  "ns -A",
	"-f":  "po -A",
	"-c":  `-o=jsonpath='{range .spec.containers[*]}{.name}{"\n"}{end}'`, // 获取原始命令 去除-c 然后获取name => k get po -n b01xm1 workspace5edc5287bd2c4f7f-647bc75db6-sxhj8 -oyaml|grep name|grep -vE 'field|-|path'
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
				// 默认基于pod
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
