import gleeunit/should

import block/op.{GT}
import block/sockets.{Sockets}

pub fn sockets_gt_test() {
  Sockets(GT, 0)
  |> sockets.to_string
  |> should.equal("Sockets > 0")
}
