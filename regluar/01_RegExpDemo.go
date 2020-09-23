package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "abc a7v  6ui avf"
	//解析编译正则表达式
	ret := regexp.MustCompile(`a[0-9]v`) //``:表示使用原生字符串
	//提取需要信息
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("alls:", alls)

	//小数匹配
	str1 := "3.14 6.15  .68 haah 1.0 abc 7 3."
	ret1 := regexp.MustCompile(`[0-9]+\.[0-9]+`)
	submatch := ret1.FindAllStringSubmatch(str1, -1)
	fmt.Println("submatch:", submatch)

	//html 内容匹配
	str2 := `<ul class="s-news-rank-content">
                                <li class="news-meta-item clearfix" data-index="0">
            <a class="title-content c-link c-font-medium c-line-ellipsis" href="https://www.baidu.com/s?cl=3&tn=baidutop10&fr=top1000&wd=%E7%BE%8E%E6%96%B9%E6%92%A4%E9%94%80%E8%B6%85%E5%8D%83%E5%90%8D%E4%B8%AD%E5%9B%BD%E5%85%AC%E6%B0%91%E7%AD%BE%E8%AF%81&rsv_idx=2&rsv_dl=fyb_n_homepage&hisfilter=1" target="_blank" >
                <span class="title-content-index c-color-gray2 top-0">1</span>
                <span class="title-content-title">美方撤销超千名中国公民签证</span>
                                                        <span class="title-content-mark mark-type-3"></span>
                            </a>
            <div class="news-count c-font-normal c-color-gray2" title="搜索指数479万">
                479万
            </div>
        </li>
                        <li class="news-meta-item clearfix" data-index="1">
            <a class="title-content c-link c-font-medium c-line-ellipsis" href="https://www.baidu.com/s?cl=3&tn=baidutop10&fr=top1000&wd=%E7%89%B9%E6%9C%97%E6%99%AE%E5%9B%9E%E5%BA%94%E6%B7%A1%E5%8C%96%E6%96%B0%E5%86%A0%E4%B8%A5%E9%87%8D%E6%80%A7&rsv_idx=2&rsv_dl=fyb_n_homepage&hisfilter=1" target="_blank" >
                <span class="title-content-index c-color-gray2 top-1">2</span>
                <span class="title-content-title">特朗普回应淡化新冠严重性</span>
                                    <span class="title-content-mark mark-type-0"></span>
                            </a>
            <div class="news-count c-font-normal c-color-gray2" title="搜索指数462万">
                462万
            </div>

			<div> fdsd </div>
			<div> hello9
            9
            9 </div>
			13918783643  15300837597
`
	ret3 := regexp.MustCompile(`<div>(.*)</div>`) //注意 如果是匹配div包括内容总含有换行符，则匹配不到47-49行数据

	submatch2 := ret3.FindAllStringSubmatch(str2, -1)
	fmt.Println("submatch2:", submatch2)

	ret4 := regexp.MustCompile(`<div>(?s:(.*?))</div>`) //注意 如果是匹配div包括内容总含有换行符，则匹配不到47-49行数据

	submatch3 := ret4.FindAllStringSubmatch(str2, -1)
	fmt.Println("submatch3:", submatch3)

	//迭代出每个成员的元素
	for _, one := range submatch3 {
		fmt.Println("one[0]=", one[0])
		fmt.Println("one[1]=", one[1])
	}

}
