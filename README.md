# monosize

File size calculator for your CLI applications for displaying more user friendly output.

Example :

```go
monosize.GetFixedSize(1024) // Output : "  1.00 KB"
monosize.GetFixedSize(1048576) // Output : "  1.00 MB"
```

**GetFixedSize** function reads files size as **float64** for supporting huge file sizes (like ZettaByte and YottaByte) and returns user friendly file size in 6 characters long with leading spaces (if required) using Base 2 calculation and file size abbreviation from Bytes to YottaBytes as **string**.

Output will be always 6+1+2 = 9 characters long until YottaByte limit is exceeded. 6 characters for file size, 1 space character and 2 characters for abbreviations.

## Installation

```sh
go get -u github.com/ozguncagri/monosize
```

## Sample Usage

```go
package main

import (
	"fmt"

	"github.com/ozguncagri/monosize"
)

func main() {
	textData := []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi erat metus, tempor in metus in, fringilla condimentum erat. Aenean felis libero, euismod et nulla nec, lobortis hendrerit arcu. Nam vulputate molestie metus, nec pellentesque lorem porta sit amet. Proin id ex est. Morbi mauris ex, elementum quis suscipit eu, scelerisque ut mi. Proin tellus tellus, commodo vel ornare ac, convallis vel felis. Curabitur venenatis arcu elit, hendrerit finibus ipsum tempus id. Phasellus ut risus et ligula sodales ultrices. Suspendisse nec bibendum ipsum, eget lacinia erat. Donec fermentum ex a nulla sagittis, vel blandit metus interdum. Aliquam luctus sapien eu metus mollis posuere. Suspendisse iaculis, diam quis cursus elementum, massa ex fringilla enim, aliquam consectetur ipsum turpis vitae enim.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vestibulum elementum elementum. Maecenas ullamcorper non urna eu sagittis. Aliquam tempor tortor magna, eu cursus est vehicula id. In finibus purus sed ligula tristique dictum. Cras elementum nisl quis facilisis vulputate. Phasellus mattis ante in tristique ultricies. Sed ante mauris, ullamcorper sit amet gravida a, faucibus vitae orci. Vestibulum libero tellus, bibendum et posuere vitae, sollicitudin et ligula. Praesent congue nisi a turpis aliquet molestie. Aliquam nec nunc nec erat porta ullamcorper nec a justo. Curabitur eleifend vestibulum molestie. Fusce elementum, justo quis aliquam imperdiet, elit sem dignissim odio, sed venenatis orci quam eu nisi. Morbi sit amet placerat elit.

Vestibulum sit amet libero nec justo molestie laoreet et at libero. Morbi fermentum quam et blandit consequat. Nullam tempus finibus mattis. Ut varius arcu eu neque condimentum, quis dignissim sem sodales. Suspendisse nec quam a est efficitur sodales suscipit eget enim. Maecenas neque ligula, dapibus ac velit eu, fringilla malesuada lacus. Ut ut diam sed augue tempor tempus. Mauris faucibus massa tellus, vel auctor risus pharetra efficitur. Fusce eu varius velit. Duis nec nunc quis tortor venenatis condimentum quis eu sem. Proin eget faucibus mi, sed vulputate nulla. Quisque eget nulla neque.

Suspendisse at libero sed purus interdum finibus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed accumsan dolor efficitur, aliquam est id, porta risus. Sed fringilla nibh quis facilisis dapibus. Sed lectus sapien, ornare nec tincidunt a, blandit vitae erat. Curabitur ac turpis sit amet eros pulvinar blandit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce tempus maximus diam, ut pretium erat condimentum a. Vivamus accumsan vel sem in consectetur. Proin laoreet lorem arcu, eu ullamcorper nisl porta molestie. Nunc ut suscipit sem, vel tempor augue. Maecenas eget metus vitae diam volutpat viverra quis at orci. Nam lacus tellus, finibus sit amet turpis in, tristique maximus magna. In hac habitasse platea dictumst.

Sed egestas nibh id ornare eleifend. Donec accumsan lectus diam, nec commodo nisl tempus sit amet. Curabitur commodo varius ipsum non imperdiet. Suspendisse interdum, ex quis commodo scelerisque, lacus nulla tristique neque, quis ultricies lectus urna non ipsum. Morbi eu ante in sem consectetur maximus nec nec lorem. Pellentesque a augue at risus tristique gravida eu ut mauris. Suspendisse vestibulum elit ante, nec convallis urna dignissim eu. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aliquam non est a dolor dapibus vestibulum. Maecenas nec justo tincidunt, vulputate tellus vel, interdum velit. Praesent cursus diam augue, dignissim sodales ipsum faucibus vel. Integer id leo felis. Curabitur eu est augue.`)

	dataSize := float64(len(textData))

	fmt.Printf("File size is : '%v'\n", monosize.GetFixedSize(dataSize))
}
```

### Output

```plain
File size is : '  3.58 KB'
```
