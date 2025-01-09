import gleeunit/should

import block/defences.{Armour, EnergyShield, Evasion}
import block/op.{Eq, GTE, LT}

pub fn armour_eq_test() {
  Armour(Eq, 100)
  |> defences.to_string
  |> should.equal("BaseArmour = 100")
}

pub fn evasion_lt_test() {
  Evasion(LT, 82)
  |> defences.to_string
  |> should.equal("BaseEvasion < 82")
}

pub fn es_gte_test() {
  EnergyShield(GTE, 120)
  |> defences.to_string
  |> should.equal("BaseEnergyShield >= 120")
}
