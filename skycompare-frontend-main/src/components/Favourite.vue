<template>
    <v-sheet class="favouriteComponent">
      <span class="cerrar" @click="exit">X</span>
      <v-card class="cardComponent">
        <v-form @submit.prevent="onSubmit">
          <v-spacer class="blind">----------------</v-spacer>
          <h3>Select your favourite airport from list below</h3>
          <v-select :items="airports" v-model="favourite"/>
          <br>

          <v-btn color="success" size="large" type="submit" variant="elevated" 
            block>
            Accept
          </v-btn>
        </v-form>
      </v-card>
       <Confirm :question="petition" :buttons="buttons" v-show="showConfirm" @close="closePetition"/>
    </v-sheet>
</template>

<script>
import { VForm, VBtn, VSheet } from "vuetify/lib";
import { VSpacer, VCard, VSelect } from "vuetify/lib";
import axios from "axios";
import Confirm from "@/components/Confirm";

export default {
  name: 'FavouriteVue',
  components: {
    VCard,
    VForm,
    VBtn,
    VSpacer,
    VSheet,
    VSelect,
    Confirm
  },
  data:()=>({
    airports: [],
    favourite: "",
    use: "",
    petition:"",
    showConfirm: false,
    buttons:[],
  }),
  mounted(){
    this.user = sessionStorage.getItem("user");
    axios.get("https://skycompare-backend.onrender.com/airports").then(res => {
      var aux = [];
      aux.push("No favourite airport");

      for(var i=0; i<res.data.length; i++){
        var airp = res.data[i];
        aux.push(airp.name + " ("+airp.IATA +")");
      }

      this.airports = aux;
    });
  },
  methods:{
    exit(){
      this.$emit("close");
    },
    async onSubmit(){
      var out = 0;
      await axios.get("https://skycompare-backend.onrender.com/favAirport", {
        params: {
          fav: this.favourite,
          user: this.user
        }
      }).then(res => {
        this.petition=res.data;
        out = 1;
      });

      if(out==1){
        this.buttons=["Close"];
        this.showConfirm=true;
      } 
      },
    closePetition(){
      this.exit();
    }
  }
}
</script>

<style scoped>
.favouriteComponent {
  position: absolute;
  width: 500px;
  height: 210px;
  border-style: solid;
  border-width: 5px;
  border-color: #ff5e00;
  border-radius: 25px;
  background: #90a0bf;
  left: 50%;
  margin-left: -250px;
  top: 50%;
  margin-top: -105px;;
}

.cardComponent{
  background: #c8ddff;
  height: 100%;
  margin-top: 0%;
  margin-left: 0px;
  max-width:100%;
  padding: 10px;
  border-radius: 20px;
}

.cerrar{
  position: absolute;
  background: "#c8ddff";
  color: #ff5e00;
  font-size: 25px;
  z-index: 100;
  width: 38px;
  padding-left: 11px;
  border-radius: 25px;
  right: 5px;
  top: 5px;
}

.cerrar:hover{
  background: #ff5e00;
  color: white;
}

.blind{
  color: #c8ddff; 
  font-size: 10px;
}

</style>
