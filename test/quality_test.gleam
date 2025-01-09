import gleeunit/should

import block/op.{NotEq}
import block/quality.{Quality}

pub fn quality_not_eq_test() {
  Quality(NotEq, 8)
  |> quality.to_string
  |> should.equal("Quality != 8")
}
