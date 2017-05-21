package download

import (
	"fmt"

	"github.com/Huawei/containerops/singular/init_config"
)

func Master_Download(ip string) {

	var fileslist = init_config.Get_files()
	for key, url := range fileslist {
		ExecCMDparams("wget", []string{"-P", "/tmp/", "-c", url})
		fmt.Printf("%s\n\n", key)
	}
	//scp -r ./config/. root@138.68.14.193:/tmp/
	LocalExecCMDparams("scp", []string{"-r", "./config/.", init_config.User + "@" + ip + ":" + "/tmp/config/"})

}
