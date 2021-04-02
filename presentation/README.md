# Let's Learn: Go, test it yourself!

![bg left:40% height:300px](https://miro.medium.com/max/1104/1*37nKRE8ZW7WWd8qb0RMi4g.png)

---

# TDW (Test Driven Workshop)

```go

func TestWorkshop(t *testing.T) {
    result := Workshop()
    if len(result.PeopleAwake) < 10 {
        t.Fatal("Make it more interesting!")
    }
    if !result.WasFun {
        t.Fatal("Make it more fun!!!11")
    }
    if len(result.PeopleLearnedSomething) == 0 {
        t.Fatal("Say sorry")
    }
}

```

---

# Overview

- Why do we test?
- What do we test?
- What's TDD?

---

# Why do we test? 
### Testing in production
![height:450px](https://i.imgur.com/mrUMURX.gif)

---

# Why do we test?
### Testing during deployment (the other TDD)
![height:450px](https://www.meme-arsenal.com/memes/d1117d98900b382d1669595ffb8b6b18.jpg)

---

# Why do we test?
![height:450px](https://i.redd.it/smihe0fwm9p01.png)

---

# What do we test?

![height:450px](https://3fxtqy18kygf3on3bu39kh93-wpengine.netdna-ssl.com/wp-content/uploads/2020/01/test-automation-pyramid.jpg)
*Benchmark-Testing

---

# What's test-driven development?

![height:500px](https://marsner.com/wp-content/uploads/test-driven-development-TDD.png)

---

## Resources

- https://github.com/stretchr/testify
- https://github.com/onsi/ginkgo  
- https://github.com/maxbrunsfeld/counterfeiter
