components {
  id: "table"
  component: "/scripts/table.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"vendors_table\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 96.0\n"
  "  y: 84.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/sprites/chicken.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
}
