package tidyUtil

import (
	"testing"
)

func BenchmarkTidyString(b *testing.B) {
	str := "<p></p><p></p><p>本文为精编版，共3754字，预计阅读10分钟</p><p><br/></p><p></p><p></p><p><strong style=\"\">作者简介</strong></p><p></p><p></p><p>梁世雄，1933生，广东省佛山市南海人。5岁丧父，幸逢恩师，入门学画，一路求学，后长期在广州美术学院任教，曾任广州美术学院教授、中国画系主任。爱人容璞是国学大师容庚之女，两人从大学起，风雨同路。育有一女，女儿、女婿结缘于广美求学时期，目前一个在广州美术学院美术馆、一个在广东省博物馆从事书画典藏研究工作，外孙女在广州美院就读。一家三代，美术之家。</p><p><br/></p><p>梁先生擅长中国山水画，为岭南画派继关山月、黎雄才之后重要的艺术大家。</p><p></p><p></p><p><br/></p><p></p><p></p><p><strong data-type=\"dy\" style=\"display:none;\">@@video=\"http://flv0.bn.netease.com/videolib1/2106/22/61ycnzecjx0/SD/movie_index.m3u8,http://flv0.bn.netease.com/videolib1/2106/22/61ycnzecjx0/SD/61ycnzecjx0-mobile.mp4\" img=\"http://videoimg.ws.126.net/cover/20210622/wmFta829L_cover.jpg\" alt=\"\" broadcast=\"in\" size=\"\" topicid=\"1000\" commentid=\"\" commentboard=\"video_bbs\" vid=\"VFC5A82DU\"@@</strong></p><p>"
	print(TidyFilter(str))
}
