# `linedesigns`

<img src="assets/recttri.png" width="400" height="400" alt="hello">

<a href="#how-to">TL;DR</a>

In elementary school I was introduced to a craft that used equally spaced dots connected by straight lines to create images. Here's the earliest version of one of these that I did as a kid:

<img src="assets/early.jpg" width="300" height="400" alt="iloveyoumom">

The bulk of these types of works I've done have been with a rule/protractor, not the messy string as above. I did a lot of these later on in years and came up with a number of different ways they could be modified make unique. Throughout my adult life I sporadically do them and when doing one recently (April 2018) I thought about writing some software to help me out in prototyping.

While this project won't be able to do the image on the left, it would neat to get it to be able to output different shapes, allow for modification of line spacing, fill (pipe dream), etc.

Thanks to [@tidwall](http://twitter.com/tidwall) for his work on [pinhole](https://github.com/tidwall/pinhole) as linedesigns is built on top of it.

<img src="assets/02.jpg" width="300" height="400" alt="iloveyoumom">
<img src="assets/07.jpg" width="300" height="400" alt="iloveyoumom">

## How To

```
$ go get -u github.com/jaysonesmith/linedesigns
```

```go
l := lines.New(0.1, 1000, 1000)
l.Triangle(0, .9, -.9, -.9, .9, -.9, 50)
l.Save("triangle.png")
```

### Output:

<img src="assets/triangle.png" width="400" height="400" alt="triangle">

## Contact

Jayson Smith [@thatengjayson](http://twitter.com/thatengjayson)

## License

`linedesigns` source code is available under the ISC [License](/LICENSE).