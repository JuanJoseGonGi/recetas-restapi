use restapi;
var bulk = db.recetas.initializeUnorderedBulkOp();
bulk.insert({
  nombre: "nombre1",
  descripcion: "descripcion1",
  ingredientes: [
    "ingrediente1",
  "ingrediente2"
]
});
bulk.insert({
  nombre: "nombre2",
  descripcion: "descripcion2",
  ingredientes: [
    "ingrediente2",
    "ingrediente2"
  ]
});
bulk.execute();
