<!-- LTeX: language=en-US -->
# GUI

<!--toc:start-->
- [GUI](#gui)
  - [Available Widgets](#available-widgets)
    - [BaseWidget](#basewidget)
    - [Label](#label)
    - [Button](#button)
    - [Scroll Panel](#scroll-panel)
  - [Internal Implementation](#internal-implementation)
    - [How to create a new Widgets](#how-to-create-a-new-widgets)
- [Available Widgets](#available-widgets)
- [How to create a new Widget](#how-to-create-a-new-widget)
<!--toc:end-->

## Usage
TODO

### Available Widgets

#### BaseWidget
This is not a real Widget. It provides core properties, and is inherited by every other Widget listed below.

**Properties**:
- `position`: A Vector2 storing the position of the Widget.
- `size`: A Vector2 storing the size of the Widget.

**Methods**:
- `SetPosition(p rl.Vector2)`
- `GetPosition() rl.Vector2`
- `SetSize(s rl.Vector2)`
- `GetSize() rl.Vector2`
- `Bounds() rl.Rectangle`

#### Label

**Definition**:
A widget used to display a piece of text.

**Usage**:
```go
label := NewLabel("Sample Text", rl.Vector2{X: 10, Y: 20}, "color:255,255,255")
```

#### Button

**Definition**:
A widget that represents an interactive button with the capability to detect hover and click states.

**Usage**:
```go
btn := NewButton("Click Me", rl.Vector2{X: 30, Y: 40}, rl.Vector2{X: 100, Y: 30}, myCallbackFunction, "color:255,0,0|bgColor:0,0,255")
```

#### Scroll Panel

**Definition**:
A scrollable container that holds and arranges other widgets inside it.

**Methods**:
- `AddChild(child Widget)`: Adds a child widget to the ScrollPanel.
- `RemoveChild(target Widget)`: Removes a specified child widget from the ScrollPanel.

**Usage**:
```go
scrollPanel := NewScrollPanel(rl.NewRectangle(10, 10, 200, 200), rl.NewRectangle(0, 0, 400, 400), "color:150,150,150")
scrollPanel.AddChild(myLabelWidget)
scrollPanel.RemoveChild(myLabelWidget)
```

## Internal Implementation
The following will describe how the `gui` package is structured and implemented.
Afterward a guide on how to implement new Widgets will follow.

The `gui` package has one main object, the `Gui`. This `Gui` object is a
container for other objects, and has a `Draw()` function. This draw function
should be called every frame while the GUI is active.

Beyond that, every widget is either atomic (such as a `Button`) or another
container (such as a `ScrollPanel`). Using container Widgets (and their `AddChild` functions), a kind of tree structure is created.

Example:
```
Gui
|- Button
|
|- Button
|
|- ScrollPanel
|   |- Label
|   |- Button
|
|- ScrollPanel
|   |- Label
|   |- Button
|
|- Label
|
```

### Walking the tree
To ensure every Widget is drawn in the correct order, the Gui-Objects `Draw()`
function actually starts a recursive drawing process, beginning at the `Gui`
node.

Here is an excerpt from that recursive function:
```go
func doRecursiveDraw(container Container) {
	for _, widget := range container.Children {
		switch w := any(widget).(type) {
		case *Label:
			w.update_label()
			backend_label(*w)
			backend_label_finalize(*w)
		case *ScrollPanel:
			w.update_scroll_panel()
			backend_scroll_panel(*w)
			doRecursiveDraw(*w.container) // draw the panels children
			backend_scroll_panel_finalize(*w)
    }
}
```

As one may see, every container Widget type (one that has its own children),
must call `doRecursiveDraw()` on its own children.

### Updating and Drawing
Each type of Widget has two main stages. The **update** and the **draw**.
While the **update** function is typically located in `pkg/gui/gui.go`, as it
is a rigid implementation of how the Widget should behave, the **draw**
function is relegated to the `pkg/gui/backend.go` file. This backend may be
swapped out with an alternative implementation.

Let's look at the `Button` Widget for example. This is an excerpt of the `doRecursiveDraw()` function:
```go
    case *Button:
        w.update_button()
        backend_button(*w)
        backend_button_finalize(*w)
```
(Here, `w` is a pointer to a `Button` widget.)

First comes the logical **update**, the `w.update_button()` function. It has a
pointer receiver, as it is able to modify the state of the button widget (which
is saved in the button struct itself, as `button.state`).

Then come to calls to the drawing backend: `backend_button(*w)` (notice that
the button is passed as a value; the draw functions will not modify state) and
then `backend_button_finalize(*w)`. These two steps may not be important for a
button, but are unavoidable in other cases.

If we look at `ScrollPanel` for example:
```go
    case *ScrollPanel:
        w.update_scroll_panel()
        backend_scroll_panel(*w)
        doRecursiveDraw(*w.container) // draw the panels children
        backend_scroll_panel_finalize(*w)
```
We can see that the `scrollPanels` children are drawn in between these steps.
This allows `backend_scroll_panel()` to begin a scissor-mode, and
`backend_scroll_panel_finalize()` to end that scissor-mode.


### How to create a new Widget
TODO

### Swapping out the backend
TODO describe contract, styledefs, etc
