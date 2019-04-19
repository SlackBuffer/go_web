- Polymorphism
- Focus
    - Commitment
    - > Meditate
- Go temporary template
    - https://play.golang.org/p/XnpQkzL1ekQ
- 标准库里的常量（变量同理）访问 - `time.Kitchen`
- Time format

    ```go
    func monthDayYear(t time.Time) string {
        // gohtml 01 的位置就是月份，02 的位置就是日期
        // gohtml 01/02 03:04:05PM '06 -0700
        // gohtml Mon Jan 2 15:04:05 MST 2006
        return t.Format("01-02-2006")
    }
    ```

- 使用标准库函数

    ```go
    func init() {
        // gohtml ParseGlob 返回的两个参数直接作为 Must 的参数
        tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
    }
    ```

- Go template comment - `{{/* I'm comment*/}}`
- Escape
    - https://golang.org/pkg/html/template/#hdr-Contexts
- [Understanding servers](https://github.com/GoesToEleven/golang-web-dev/tree/master/014_understanding-servers)
- Go packages - https://godoc.org/
- net/http 处理以 `/` 结尾的路由，会匹配子路由
- 网页的图片会发起单独的 http 请求，需要显示地 serve 对应文件
- `/` 开头的绝对路径相对于 `index.html` 的位置