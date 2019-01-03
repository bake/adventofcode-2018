# Advent of Code 2018

Solutions for [Advent of Code 2018](https://adventofcode.com/2018/) written in Go.

## Some interesting Plots

|     **Day 6: Chronal Coordinates**     |     **Day 11: Chronal Charge**     |
| :------------------------------------: | :--------------------------------: |
|   ![Day 6](/day-06/part-01/grid.png)   | ![Day 11](/day-11/part-01/out.png) |
| **Day 18: Settlers of The North Pole** |                                    |
|   ![Day 18](/day-18/part-01/out.gif)   |                                    |

### Animation

Plot each generation (usually with an `Image() *image.RGBA` function) and use ffmpeg to gule them together. See [day 18](/day-18/part-01/main.go) for an example.

```bash
$ ffmpeg -framerate 30 -i out/%d.png -vf "scale=iw*5:-1" -sws_flags neighbor -y out.mp4
```
