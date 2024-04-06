import gleam/result
import simplifile

pub fn read_text(file_path: String) -> Result(String, String) {
  simplifile.read(file_path)
  |> result.map_error(simplifile.describe_error)
}

pub fn write_text(data: String, file_path: String) -> Result(Nil, String) {
  simplifile.write(file_path, data)
  |> result.map_error(simplifile.describe_error)
}
