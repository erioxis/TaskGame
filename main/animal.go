components {
  id: "animal"
  component: "/scripts/animal.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"chicken\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/sprites/chicken.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
embedded_components {
  id: "pet_sound"
  type: "sound"
  data: "sound: \"/sounds/pet.wav\"\n"
  ""
}
embedded_components {
  id: "pop_sound"
  type: "sound"
  data: "sound: \"/sounds/pop.wav\"\n"
  ""
}
