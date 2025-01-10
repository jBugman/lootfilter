import gleeunit/should

import poe/actions
import poe/actions/label/border_color.{BorderColor}
import poe/actions/label/rgba.{Greyscale}
import poe/block
import poe/conditions
import poe/conditions/class
import poe/conditions/rarity
import poe/op

pub fn block_show_test() {
  block.Show([conditions.Sockets(op.GT, 0)], [
    actions.BorderColor(BorderColor(Greyscale(127))),
  ])
  |> block.to_string
  |> should.equal(
    "Show
  Sockets > 0
  SetBorderColor 127 127 127 255
  EnableDropSound",
  )
}

pub fn block_show_with_sound_test() {
  block.Show([conditions.Sockets(op.GT, 0)], [
    actions.BorderColor(BorderColor(Greyscale(127))),
    actions.Alert(16, 200),
  ])
  |> block.to_string
  |> should.equal(
    "Show
  Sockets > 0
  SetBorderColor 127 127 127 255
  PlayAlertSound 16 200
  DisableDropSoundIfAlertSound",
  )
}

pub fn block_hide_test() {
  block.Hide([
    conditions.Class([class.Quarterstaves]),
    conditions.Rarity(op.Eq, rarity.Normal),
  ])
  |> block.to_string
  |> should.equal(
    "Hide
  Class \"Quarterstaves\"
  Rarity = Normal
  SetFontSize 20
  DisableDropSound",
  )
}
