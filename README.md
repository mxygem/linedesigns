# `linedesigns`
<img src="assets/trioct.png" width="300" height="300" alt="trioct"><img src="assets/mixed.png" width="300" height="300" alt="mixed">
<img src="assets/freeform.png" width="300" height="300" alt="freeform"><img src="assets/recttri.png" width="300" height="300" alt="recttri">

`linedesigns` is the digitization of an arts craft I was taught in my elementary school days. It's meant to help prototype new designs quickly since doing these designs by hand can take hours.

---
## Install
```
$ go get -u github.com/jaysonesmith/linedesigns
```
---
## How To
`linedesigns` commands generally follow the pattern of taking in a series of coordinates that process/connect to one another in a counter-clockwise fashion.

Each run consists of three main steps: 
- **Object Creation** - setup a `linedesign` object to use by defining default line thickness, image width, and image height.
- **Design Creation** - These commands are what make the actual designs. You can perform as many design creation steps as you'd like in order to add multiple designs to a single image. Coordinate points are based off a Cartesian plane with (0,0) as the centered origin point and 1 being the max displayable point on either axis.
- **Saving** - Finalizes the current `linedesign` object and writes the file to disk.

### Start
The following examples will all assume that a `linedesigns` object has been created with the following:
```go
// new input: line width, dot size, image width, image height
l := lines.New(0.1, 0.05, 1000, 1000)
```


### Design Creators
`linedesign's` creators are split into two main groups based on their outputs:
- **Dotted** - Outputs only dots at each point position instead of connecting each by lines. This output can be used if you'd like to connect the lines yourself for instance. Takes in N number of coordinates to create designs from, a count of how many points per line segment/design to create, and a dot size.
- **Lined** - Outputs designs connected by lines. Currently the only method for line connection is via straight & black lines. These creators take in N amount of coordinates based on the shape to be created unless otherwise noted below.

#### Dotted Creators

**LineDotted**

The simplest command available to `linedesign`. It takes in two coordinates, a count, and outputs a single, dotted line.
```go
l.LineDotted(-.8, 0, .8, 0, 25)
```
<img src="assets/linedotted_ex.png" width="300" height="300" alt="linedotted_ex">

**AngleDotted**

One step up from a single line. It takes in three coordinates, a count, and outputs a single, dotted angle.
```go
l.AngleDotted(-.6, .6, 0, -.6, .6, .6, 25)
```
<img src="assets/angledotted_ex.png" width="300" height="300" alt="angledotted_ex">

**TriangleDotted**

TriangleDotted takes in three coordinates, a count, and outputs a dotted triangle.
```go
l.TriangleDotted(0, .9, -.9, -.9, .9, -.9, 50)
```
<img src="assets/triangledotted_ex.png" width="300" height="300" alt="triangledotted_ex">

**RectangleDotted**

RectangleDotted takes in four coordinates, a count, and outputs a dotted rectangle.
```go
l.RectangleDotted(-.8, .8, -.8, -.8, .8, -.8, .8, .8, 25)
```
<img src="assets/rectangledotted_ex.png" width="300" height="300" alt="rectangledotted_ex">

**CircleDotted**

CircleDotted takes in a center point, a radius, a count, and outputs a dotted circle.
```go
l.CircleDotted(0, 0, .8, 100)
```
<img src="assets/circledotted_ex.png" width="300" height="300" alt="circledotted_ex">

**FreeformDotted**

FreeformDotted takes in a collection(slice of slices) of three or more coordinates, a count, and outputs a dotted whateveryoucreated. *Note: The two Freeform commands connect the last provided coordinate to the first!*
```go
coords := [][]float64{
    []float64{-.9, .9},
    []float64{-.2, .0},
    []float64{-.9, -.9},
    []float64{0, -.9},
    []float64{0, 0},
    []float64{.9, 0},
    []float64{.9, .9},
}
l.FreeformDotted(coords, 50)
```
<img src="assets/freeformdotted_ex.png" width="300" height="300" alt="freeformdotted_ex">

#### Lined Creators

**Angle**

The simplest of the lined creators. It takes in three coordinates, a count, and outputs a single, lined angle.
```go
l.Angle(-.6, .6, 0, -.6, .6, .6, 25)
```
<img src="assets/angle_ex.png" width="300" height="300" alt="angle_ex">

**Triangle**

Triangle takes in three coordinates, a count, and outputs a lined triangle. Similar to `Angle`, except this connects each combo of line segments. (1 & 2, 2 & 3, 3 & 1)
```go
l.Triangle(0, .9, -.9, -.9, .9, -.9, 50)
```
<img src="assets/triangle_ex.png" width="300" height="300" alt="triangle_ex">

**Rectangle**

Rectangle takes in four coordinates, a count, and outputs a lined rectangle.
```go
l.Rectangle(-.6, .6, -.6, -.6, .6, -.6, .6, .6, 25)
```
<img src="assets/rectangle_ex.png" width="300" height="300" alt="rectangle_ex">

**Circle**

Circle takes in a center point, a radius, a count, an offset, a wrap bool, and outputs a lined circle.
- Offset is a number that defines how many points `linedesign` will skip to connect two points with a line. This can have neat outputs based on what you input. 
- Wrap tells `linedesign` whether to continue around through to make sure all points created are connected. Without it, depending on your offset value, you may not have all points connected which can also lead to other neat designs.

*The below examples show the same parameters but with wrap on and off for each photo respectively.
```go
l.Circle(0, 0, .8, 100, 40, true)
```
<img src="assets/circle_ex.png" width="300" height="300" alt="circle_ex"><img src="assets/circle_nowrap_ex.png" width="300" height="300" alt="circle_nowrap_ex">

**Freeform**

Freeform takes in a collection(slice of slices) of three or more coordinates, a count, an offset, a wrap bool, and outputs a lined whateveryoucreated.
```go
coords := [][]float64{
    []float64{-.9, .9},
    []float64{-.2, .0},
    []float64{-.9, -.9},
    []float64{0, -.9},
    []float64{0, 0},
    []float64{.9, 0},
    []float64{.9, .9},
    []float64{-.9, .9},
    []float64{-.2, .0},
}
l.Freeform(coords, 50, 51, false)
```
<img src="assets/freeform_ex.png" width="300" height="300" alt="freeform_ex">


---
## Color

Eventually, I plan to utilize Pinhole's color options and allow for the lines drawn to be other than black, but that is still yet to come. Additionally, I'd like to implement a way to automate the fill patterns below as well as others, but that'll be quite the challenge. (PRs welcomed!)

<img src="assets/aquax.png" width="300" height="300" alt="aquax"><img src="assets/circle_colored_colorful.png" width="300" height="300" alt="circle_colored_colorful">


## Background

Line designs are essentially a series of points connected by straight or curved lines to create neat looking designs. To further add to these designs, the squares created by the crossing of lines can be filled in different patterns. (as shown above)

If I remember correctly, I was introduced to this type of art in 4th or 5th grade math class. Here's one of the earliest versions of one of these that I did as a kid:

<img src="assets/early.jpg" width="300" height="400" alt="iloveyoumom">

That one had a piece of graph paper taped to the back side that was used as a guide for sewing the thread through on the other side.

The bulk of these types of works I've done have been with a rule/protractor, not the messy string as above. I did a lot of these later on in years and came up with a number of different ways they could be modified to be made more unique. Throughout my adult life I've sporadically done them and when doing one recently (April 2018) I thought about writing some software to help me out in prototyping.

Some of my older drawings:

<img src="assets/02.jpg" width="300" height="400" alt="2002"><img src="assets/07.jpg" width="300" height="400" alt="2007">

## Thanks

Thanks to [@tidwall](http://twitter.com/tidwall) for his work on [pinhole](https://github.com/tidwall/pinhole) as linedesigns is built on top of it and this readme has been heavily influenced by [pinhole's](https://github.com/tidwall/pinhole).

## Contact

Jayson Smith [@thatengjayson](http://twitter.com/thatengjayson) or open a [github issue.](https://github.com/jaysonesmith/linedesigns/issues)

## License

`linedesigns` source code is available under the ISC [License](/LICENSE).