const app = new Vue({
  el: '#app',
  data: {
    nombre: '',
    recetas: [],
    cargando: false,
    errado: false,
    crear: {
      nombre: '',
      descripcion: '',
      ingredientes: [],
      pasos: [],
      _id: '',
    },
    editar: {
      nombre: '',
      descripcion: '',
      ingredientes: [],
      pasos: [],
      _id: '',
    },
  },
  watch: {
    nombre: function() {
      this.getRecetas().catch((error) => {
        console.error(error);
        this.errado = true;
      });
    },
  },
  methods: {
    crearReceta: async function() {
      const response = await axios.post(
        'http://localhost:8000/recetas',
        this.crear,
      );
      this.getRecetas();
      return response.data;
    },
    addPasoCrear: function() {
      this.crear.pasos.push('');
    },
    addIngrCrear: function() {
      this.crear.ingredientes.push('');
    },
    editarReceta: async function() {
      const response = await axios.put(
        `http://localhost:8000/recetas`,
        this.editar,
      );
      this.getRecetas();
      return response.data;
    },
    addPasoEditar: function() {
      this.editar.pasos.push('');
    },
    addIngrEditar: function() {
      this.editar.ingredientes.push('');
    },
    getRecetas: async function() {
      this.cargando = true;
      const response = await axios.get(
        `http://localhost:8000/recetas/${this.nombre || ''}`,
      );
      this.recetas = response.data;
      this.cargando = false;
      return response.data;
    },
    fillEdit: function(receta) {
      for (const key in receta) {
        if (receta.hasOwnProperty(key)) {
          this.editar[key] = receta[key];
        }
      }
    },
    delReceta: async function(receta) {
      const response = await axios.delete(
        `http://localhost:8000/recetas/${receta.nombre}`,
      );
      this.getRecetas();
      return response.data;
    },
  },
  beforeMount: function() {
    this.getRecetas();
  },
});
