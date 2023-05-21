![mule validation library](cover.png)

The simple validation library for Go.

## 🚀 Installation

```shell
go get -u github.com/arivictor/mule
```

## ⚡️ Quick Start
Mule is ready to-go out of the box:

```go
m := mule.New()

m.Check(1 == 2-1, "equal", "must be equal")

if ok := m.Valid(); !ok {
    fmt.Println("validation failed")
}
```

but comes with a lot of features! Here's an example:

```go
	m := mule.New()

	m.Check(m.In("a", "a", "b", "c"), "in", "'a' must be in list")
	m.Check(m.Matches("x", regexp.MustCompile("xXx")), "match", "string must match")
	m.Check(m.Unique([]string{"a", "a", "b"}), "unique", "slice must be unique")

	if ok := m.Valid(); !ok {
		fmt.Println("validation failed")
	}
```

# ✨ Contributing

Contributions to this project are welcome. If you encounter any issues or have suggestions for improvements, please open an issue on the GitHub repository.

---

This project is licensed under the [GNU v3 License](LICENSE). Feel free to use and modify the code as per the terms of the license.


# ❤️ Acknowledgement

Based on the book "Let's Go Further" by Alex Edwards.