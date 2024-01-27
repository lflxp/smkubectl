package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

// // 解析命令
// func parseCmd(in string) {
// 	tmp := strings.Split(strings.TrimSpace(in), " ")
// 	// fmt.Printf("in %s %v %d\n", in, tmp, len(tmp))
// 	switch tmp[0] {
// 	case "kubectl":
// 		if len(tmp) == 1 {
// 			// os.Stdout.Write([]byte("get describe edit set"))
// 			for _, x := range strings.Split("get describe edit set", " ") {
// 				fmt.Println(x)
// 			}
// 			return
// 		} else if len(tmp) > 1 {
// 			switch tmp[len(tmp)-1] {
// 			case "-n":
// 				for _, x := range strings.Split("default kube-system eclipse-che", " ") {
// 					fmt.Println(x)
// 				}
// 			case "describe", "edit", "get":
// 				for _, x := range strings.Split("po pod ns namespace deploy job ingress svc cm secret", " ") {
// 					fmt.Println(x)
// 				}
// 			}
// 		}
// 		if tmp[len(tmp)-1] == "-n" {
// 			execCmd("ps -ef")
// 		}
// 	default:
// 		execCmd(in)
// 	}
// }

// func execCmd(in string) error {
// 	cmd := exec.Command("sh", "-c", in)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout // 标准输出
// 	cmd.Stderr = os.Stderr // 标准错误
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 		return err
// 	}
// 	return nil
// }

// func main() {
// 	// for idx, args := range os.Args {
// 	// 	fmt.Println("参数 "+strconv.Itoa(idx)+":", args)
// 	// }
// 	// cmd := exec.Command("sh", "-c", fmt.Sprintf("%s | sed 1d | FZF_DEFAULT_OPTS=\"--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS\" fzf -m", strings.Join(os.Args[1:], " ")))
// 	// cmd := exec.Command("sh", "-c", strings.Join(os.Args[1:], " "))
// 	parseCmd(strings.Join(os.Args[1:], " "))
// }
