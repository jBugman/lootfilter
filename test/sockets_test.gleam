import gleeunit/should

import poe/conditions/sockets.{Sockets}
import poe/op.{GT}

pub fn sockets_gt_test() {
  Sockets(GT, 0)
  |> sockets.to_string
  |> should.equal("Sockets > 0")
}
