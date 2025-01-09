import gleeunit/should

import block/base_types
import block/conditions

pub fn conditions_base_type_test() {
  base_types.BaseTypes(["Short Sword", "Short Bow"])
  |> conditions.BaseType
  |> conditions.to_string
  |> should.equal("BaseType \"Short Sword\" \"Short Bow\"")
}
