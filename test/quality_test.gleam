import gleeunit/should

import poe/conditions/quality.{Quality}
import poe/op.{NotEq}

pub fn quality_not_eq_test() {
  Quality(NotEq, 8)
  |> quality.to_string
  |> should.equal("Quality != 8")
}
