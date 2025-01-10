import gleeunit/should

import poe/conditions/class.{Class, Classes}

pub fn class_test() {
  Class("Bows")
  |> class.to_string
  |> should.equal("Class \"Bows\"")
}

pub fn classes_test() {
  Classes(["Bows", "Wands"])
  |> class.to_string
  |> should.equal("Class \"Bows\" \"Wands\"")
}
