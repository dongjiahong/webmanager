package ffmpeg

import (
	"fmt"
	"os/exec"
	"strings"

	"webmanager/util"
)

// 给要拼接的视频名，返回拼接的视频，或者错误信息
func JoinVideo(args string) (string, error) {
	fileName := make([]string, 0, 2)
	var outputName string
	fmt.Println(" join ===> name: ", args)
	for _, name := range strings.Split(args, ",") {
		fileName = append(fileName, fmt.Sprintf("file '%s'", util.GetCommonPath("ori")+name))
		outputName += name
	}

	if len(fileName) == 0 {
		return "", fmt.Errorf("no video give")
	}

	fileContent := strings.Join(fileName, "\n")
	tmpFile := util.GetCommonPath("tmp") + "tmp"
	if err := util.WriteFile(tmpFile, []byte(fileContent)); err != nil {
		return "", fmt.Errorf("write tmp file err: %v", err)
	}
	// ffmpeg -f concat -i filelist.txt -c copy output

	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", tmpFile, "-c", "copy", util.GetCommonPath("out")+outputName+".mp4")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Execute command failed: ", err, " out: ", string(output))
		return "", fmt.Errorf("%v output: %s", err, string(output))
	}
	return outputName, nil
}

func mainx() {
	f, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("Execute lookpath failed: ", err)
		return
	}
	fmt.Println(" path: ", f)

	cmd := exec.Command("ffmpeg", "-y", "-f", "gif", "-i", "af4fd8d158173152c9dfd87f5881a348.gif", "-vf", "pad=600:244:100:100", "pad.mp4")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Execute command failed: ", err, " out: ", string(output))
		return
	}
	fmt.Println("output: ", string(output))
}