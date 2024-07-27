package main

import (
	"golang.org/x/sync/errgroup"
	"io"
	"k8s.io/klog/v2"
	"net/http"
)

func main() {
	eg := errgroup.Group{}

	urls := []string{
		//"http://www.google.com",
		"http://www.baidu.com",
		"http://www.facebook.com",
	}
	//_ = urls

	for k, url1 := range urls {
		klog.Infoln(url1)
		k, url1 := k, url1
		eg.Go(func() error {
			klog.Infoln(k, url1)
			resp, err := http.Get(url1)
			//defer resp.Body.Close()
			if err != nil {
				klog.Infoln(err.Error())
				return err
			} else {
				//klog.Infoln(resp.Body)
				content, err := io.ReadAll(resp.Body)
				if err != nil {
					//klog.Infoln(err.Error())
					return err
				}
				klog.Infoln(string(content))
				defer resp.Body.Close()
				return nil
			}
		})
	}

	err := eg.Wait()
	if err != nil {
		klog.Infoln(err.Error())
	} else {
		klog.Infoln("successfully !")
	}
}
