# What are variables?

Variables are like boxes - they store things. 

Each variable has a location where it is saved. To be remembered for things like getting it later; taking things out; putting them in; even destroying it.

<img height="300" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fd2xjmi1k71iy2m.cloudfront.net%2Fdairyfarm%2Fid%2Fimages%2F443%2F0744375_PE743278_S5.jpg&f=1&nofb=1">

Simple expressions (a , b) can be combined into larger expressions with operators such as addition and subtraction (a + b). These larger expressions, sometimes called basic types, can be collected into [aggregates](https://www.computerhope.com/jargon/a/aggregat.htm) like arrays and structs. Arrays in python are Lists. And Structs in python are Dictionaries.

[Expressions](https://www.computerhope.com/jargon/e/expressi.htm) are used within a single line of code, called a [statement](https://www.computerhope.com/jargon/s/statemen.htm). A statement carries out a specific task. And the order that the task is carried out is determined by [control flows](https://www.computerhope.com/jargon/c/contflow.htm) like *if* and *for*. Statements can also be grouped into [functions](https://www.computerhope.com/jargon/f/function.htm) to be separated from the main code and/or used again.

---

# Variables in Go

Naming of Go functions, variables, constants, types, statement labels, and packages begins with a [letter](https://www.computerhope.com/jargon/l/letter.htm) or an under_score. (sidenote: *If you want to see how different letters are recogized by computers click the link before and choose any letter.*) Naming in Go is [case-sensitive](https://www.computerhope.com/jargon/c/casesens.htm): computer and Computer are different names.

There are words that should not be used for naming things in Go. These words already have set meanings in Go - used for [syntax](https://www.computerhope.com/jargon/s/syntax.htm):

<details>
  <summary>Click to see words!</summary>
  break<br>case<br>chan<br>const<br>continue<br>default<br>defer<br>else<br>fallthrough<br>for<br>func<br>go<br>goto<br>if<br>import<br>interface<br>map<br>package<br>range<br>return<br>select<br>struct<br>switch<br>type<br>var
</details>
<br>These words are not reserved in Go, but they have certain meanings within Constants, Types, and Functions so be careful when using them. Things could get confusing for everyone (computer and reader of the code)
<details>
  <summary>Constant</summary>
  true<br>false<br>iota<br>nil
</details>
<details>
  <summary>Type</summary>
  int<br>int8<br>int16<br>int32<br>int64<br>uint<br>unit8<br>uint16<br>uint32<br>uint64<br>uintpr<br>float32<br>float64<br>complex128<br>complex64<br>bool<br>byte<br>rune<br>string<br>error
</details>
<details>
  <summary>Function</summary>
  make<br>len<br>cap<br>new<br>append<br>copy<br>close<br>delete<br>compex<br>real<br>imag<br>panix<br>recover
</details>
<br>

If a variable is [declared](https://www.computerhope.com/jargon/d/declare.htm) within a function it is a local only to that function. Almost like citizenship. If you are born say in America you only have a recognized citizenship in America. You cannot be a recognized citizen say for Kenya. For instance:

```go
func america() {
    citizen = "Austin"
}

func kenya() {
    fmt.Println(citizen)
}

func main() {
    kenya()     //will return an error because kenya doesn't see the variable citizen which contains a string "Austin"
}

```
If the variable is declared outside a function; however, it can be seen and used by all files of the package to which it belongs. This is confusing and will get to it later. The case of the first letter of the name determines the it's visibility across package boundaries. If the name starts with an upper case `Printf` or `Println` , it means it's *exported* (visible everywhere outside of its own package and may be used by other parts of the program). I wouldn't worry about this yet.

Just know names declared in Go have boundaries. Those boundaries are pre-set by the first letter being capital or lower-case.

Go programmers use camel-case to form names for things like variables and functions.
`prefered: camelCase or CamelCase`
`not-prefered: camel_case, camel_Case`
The letters of acronyms such as HTML or ASCII should be named in functions or variables like `htmlEscape, HTMLEscape, escapeHTML` **not** `escapeHtml` because they're read by the computer in all 