<template>
<div>
                                   <div style="padding-top= 20px" >
                        <div class="gridizquierdo">
                            <ul>
                                <li>Ubicación: </li>
                                <li>Capacidad de la sala:</li>
                                <li>Cantidad de asientos: </li>
                            </ul>
                        </div>
                        <div class="gridderecho">
                            <ul>
                                <li >{{this.salasdetalle[0].noasientos}}</li>
                                <li v-text="asientos"></li>
                                <li > </li>
                            </ul>
                        </div>
                </div>
</div>
    
</template>

<script>
export default {
props: [
    "salaElejida"
],
data(){
    return {
                
                sala: [],
                equipo: [],
                mobiliario: [],
    }
},
mounted(){
    this.regreso()    
},
methods : {
            
            regreso(){
                axios.get(`/api/salas/${this.salaElejida}`).then((response) => {this.sala = response.data;})
                this.$emit('numeroasientos', this.sala[0])
            },
            consultasal(){
                for (let i in this.salas) {
                   let asientos = this.salas[i].noasientos;
                   if (this.asistentes <= asientos) {
                   }
                   else {this.salas.splice(i,1);}
                   asientos = null;
                }
            },
            salaDetalles(){
                return console.log(this.sala[0].nosientos)
            }
}
}
</script>