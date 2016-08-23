package main

import "http"

func main() {
  http.Parse(
    "GET /tutorials/other/top-20-mysql-best-practices/ HTTP/1.1\n" +
    "Host: net.tutsplus.com\n" +
    "User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)\n" +
    "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\n" +
    "Accept-Language: en-us,en;q=0.5\n" +
    "Accept-Encoding: gzip,deflate\n" +
    "Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\n" +
    "Keep-Alive: 300\n" +
    "Connection: keep-alive\n" +
    "Cookie: PHPSESSID=r2t5uvjq435r4q7ib3vtdjq120\n" +
    "Pragma: no-cache\n" +
    "Cache-Control: no-cache\n" +
    "NOPE\n")
}
