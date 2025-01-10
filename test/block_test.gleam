import gleeunit/should

import poe/actions
import poe/actions/label/border_color.{BorderColor}
import poe/actions/label/rgba.{Greyscale}
import poe/block.{Show}
import poe/conditions
import poe/op.{GT}

pub fn block_show_test() {
  Show([conditions.Sockets(GT, 0)], [
    actions.BorderColor(BorderColor(Greyscale(127))),
  ])
  |> block.to_string
  |> should.equal(
    "Show
  Sockets > 0
  SetBorderColor 127 127 127 255",
  )
}
