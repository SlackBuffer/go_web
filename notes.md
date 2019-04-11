- Polymorphism
- Focus
    - Commitment
    - > Meditate
- Go temporary template
    - https://play.golang.org/p/XnpQkzL1ekQ
- 标准库里的常量访问 - `time.Kitchen`
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