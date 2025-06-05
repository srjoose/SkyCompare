<template>
  <div class="Main">

    <div class="menu">
      <img alt="Vue logo" src="../assets/logo.png" height="100%" width="5%" class="logoMain" @click="goBack">
      <span @click="goBack">SkyCompare®</span>
      <nav class="navigator">
        <v-btn class="vbtn" :disabled="!user" @click="saveComparison">Save comparison</v-btn>
        <v-btn class="vbtn" :disabled="!user" @click="showFavourite=true">Set favourite airport</v-btn>
      </nav>
      <p id="login" @click="login" v-show="showLogin">Login</p>
      <p class="user" v-show="!showLogin" @click="closeSession">{{ nickname }}</p>
    </div>
    <v-card class="selects">
      <v-select v-model="departure" :items="airportsDep" class="airfield" label="Departure Airport" @change="updateAirport('departures')"></v-select>
      <v-btn class="goBtn" @click="go">GO</v-btn>
      <v-select v-model="arrival" :items="airportsArr" class="airfield" label="Arrival Airport" @change="updateAirport('arrivals')"></v-select>
      <v-text-field 
        label="Date" 
        readonly 
        prepend-icon="mdi-calendar" 
        v-model="date" 
        @click="calendar"/>
      </v-card>

      <v-date-picker 
        v-show="showCalendar" 
        class="calendar" 
        v-model="date"
        @change="selectDate"/>

    <div class="comparator" v-show="showFly">
      <v-card class="fly" v-for="(fly, count) in flies" :key="count" :index="count">
        <h3>{{ departure }}&nbsp;&nbsp;&nbsp;►►►►&nbsp;&nbsp;&nbsp;{{ arrival }}</h3>

        <div class="company">
          <h1 class="companyH1C"> {{ fly.company.name }}</h1>
          <h2 class="time">{{ fly.timeDep }} → {{ fly.timeArr }}</h2>
          <h1 class="companyH1P"> {{ fly.price }}€</h1>
          <img src="../assets/oferta.png" v-show="fly.sales==1" class="sales"/>
          <v-btn class="companyVBtn">Generate ticket</v-btn>
        </div>
      </v-card>
    </div>
    <Confirm :question="peticion" :buttons="buttons" v-show="showConfirm" @close="closePetition"/>
    <Favourite v-show="showFavourite" @close="closeFavourite"/>
  </div>
</template>

<script>
import Confirm from "@/components/Confirm.vue";
import Favourite from "@/components/Favourite.vue"; 
import {VBtn, VCard, VSelect,VDatePicker, VTextField} from "vuetify/lib";
import axios from "axios";

export default {
  name: 'MainComponent',
  components:{
    Confirm,
    Favourite,
    VBtn,
    VCard,
    VSelect,
    VDatePicker,
    VTextField,
  },  
  data:()=>({
    showLogin: true,
    showConfirm: false,
    showFly: false,
    showCalendar: false,
    showFavourite: false,
    nickname: "",
    peticion: "",
    airports: [],
    airportsDep: [],
    airportsArr: [],
    flies: [],
    buttons:[],
    user: false,
    departure: "",
    arrival:"",
    date: null,
  }),
  mounted(){
    const user = sessionStorage.getItem("user");
    const date = new Date();
    this.date = date.toLocaleDateString();
    if(user!=null && user!=""){
      this.nickname = user;
      this.showLogin=false;
      this.user = true;
    }
    this.getAirports();

    if(user!=null){
      this.getFavouriteAirport();
    }
  },  
  methods:{
    goBack(){
      this.$router.push("/");
    },  
    login(){
      this.$router.push("/Login");
    },
    closeSession(){
      this.peticion="Do you really want close your session";
      this.buttons =["Yes", "No"]
      this.showConfirm=true;
      this.departure="";
      this.arrival="";
      this.flies=[];
    },
    closePetition(code){
      this.showConfirm=false;
      if(code=="Yes"){
        this.showLogin=true;
        sessionStorage.removeItem("user");
        this.user=false;
      }
    },
    getAirports(){
      var airp = [];
      var airpD = [];
      var airpA = [];
      axios.post("https://skycompare.onrender.com/airports").then(r=>{
        for(var i=0; i<r.data.length; i++){
          airp.push(r.data[i].name + " ("+r.data[i].IATA +")");
          airpD.push(r.data[i].name + " ("+r.data[i].IATA +")");
          airpA.push(r.data[i].name + " ("+r.data[i].IATA +")");
        }
        this.airports = airp;
        this.airportsArr=airpA;
        this.airportsDep=airpD;
      });
    },
    async updateAirport(type){
      var airp = [];
      var airport = this.arrival;
      this.showFly = false;
      if(type=="departures") airport = this.departure;
      var IATA = airport.substring(airport.indexOf("(")+1, airport.indexOf(")"));

      await axios.get("https://skycompare.onrender.com/airports/selected", {
        params:{
          IATA: IATA,
        }
      }).then(r=>{
        for(var i=0; i<r.data.length; i++){
          airp.push(r.data[i].name + " ("+r.data[i].IATA +")");
        }
        console.log(airp);
        if(type=="departures"){
          this.airportsArr = airp;
        }else{
          this.airportsDep = airp;
        }
      });
    },

    async go(){
      var departure = this.departure.substring(this.departure.indexOf("(")+1, this.departure.indexOf(")"));
      var arrival = this.arrival.substring(this.arrival.indexOf("(")+1, this.arrival.indexOf(")"));

      await axios.get("https://skycompare.onrender.com/routes", {
        params:{
          dep: departure,
          arr: arrival
        }
      }).then(res =>{
        if(res.data[0].route.price==0){
          this.peticion="Sorry, we did not find available flies for this route on that dates.";
          this.buttons=["Close"]
          this.showConfirm=true;
        }else{
          this.showFly = true;
          for(var i=0; i<res.data.length; i++){
            res.data[i].price = Math.round(res.data[i].price, -2);
          }        
          
          res.data.sort(function (a, b) {
            if (a.price > b.price) {
              return 1;
            }
            if (a.price < b.price) {
              return -1;
            }
            // a must be equal to b
            return 0;
          });
          this.flies = res.data;
        }
      });
    },
    calendar(){
      this.showCalendar=true;
    },
    selectDate(){
      this.showCalendar=false;
    },
    closeFavourite(){
      this.showFavourite=false;
      this.getFavouriteAirport();
    },
    getFavouriteAirport(){
      axios.get("https://skycompare.onrender.com/favAirportGet", {
        params:{
          nickName: this.nickname
        }
      }).then(res =>{
        console.log(res.data);
        this.departure = res.data.airport;
      });
    },
    saveComparison(){

      console.log(this.flies);
      var save = "";
      var companies="";
      var prices=""
      var hoursDep="";
      var hoursArr="";
      for(var i=0; i<5; i++){
        if(i<this.flies.length){
          companies = companies + this.flies[i].company.id+" ";
          hoursDep = hoursDep + this.flies[i].timeDep+" ";
          hoursArr = hoursArr + this.flies[i].timeArr + " ";
          prices = prices + this.flies[i].price + " ";
        }else{
          companies = companies + this.flies[0].company.id+" ";
          hoursDep = hoursDep + this.flies[0].timeDep+" ";
          hoursArr = hoursArr + this.flies[0].timeArr + " ";
          prices = prices + this.flies[0].price + " ";
        }
      }

      axios.get("https://skycompare.onrender.com/save", {
        params:{
          nickname: this.nickname,
          route_id: this.flies[0].route.id,
          companies: companies,
          prices: prices,
          hoursDep: hoursDep,
          hoursArr: hoursArr,
          flies: save,
        }
      }).then(res => {
        alert(res.data);
      });
    }
  }
}
</script>

<style scoped>
.Main{
  background-image: url("../assets/image1.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  position:absolute;
  top: 0vh;
  left: 0vw;
  width: 100vw;
  height: 100vh;
}

.menu{
  position: absolute;
  margin-top: 0vh;
  color: white;
  background-color: #90a0bf;
  width: 100%;
  height: 7.2vh;
  font-size: 60px;
  display: flex;
  padding-left: 20%;
  font-size: 10px;
}

span{
  position: absolute;
  left: 7vw;
  top: 0vh;
  font-size: 25px;
  font-weight: bold;
}

p{
  position:absolute;
  right: 10vw;
  margin-top: 0.5vh;
  font-size: 25px;
}

p:hover{
  color: #ff5e00;
  border-style: solid;
  border-radius: 10px;
  border-color: #ff5e00;
  margin-top: 0.3vh;
  right: 9.85vw;
}

.logoMain{
  margin-top: 0.5vh;
  margin-left: -18vw;
}

.button1{
  height: 100px;
  width: 200px;
  font-size: 40px;
  background-color: black;
  color: cyan;
  border-style: solid;
  border-width: 6px;
  border-color: cyan;
  border-radius: 30px;
}

.navigator{
  margin-left: 30vw;
  padding-top: 5px;
}

.vbtn{
  margin-left: 10px;
  border-style: solid;
  border-color: #ff5e00;
  border-radius: 15px;
}

.selects, .fly{
  margin-top: 60px;
  display: flex;
  width: 98vw;
  margin-left: 0.5vw;
  border-style: solid;
  border-color: #ff5e00;
  border-radius: 15px;
  border-width: 3px;
}

.airfield{
  margin-left: 8vw;
  padding-right: 8vw;
  max-width: 490px;
}

.route{
  display: block;
}

.goBtn{
  margin-top: 15px;
  border-style: solid;
  border-color: #ff5e00;
  border-radius: 15px;
}

.fly{
  margin-top: 30px;
  width: 99%;
  padding-left: 20px;
  display: block;
}

.company{
  display: flex;
  width: 100vw;
}

.companyH1P{
  position: absolute;
  width: 100%;
  text-align: center;
}

.companyVBtn{
  position: absolute;
  right: 10px;
  height: 40px;
  border-style: solid;
  border-color: #ff5e00;
  border-radius: 10px;
  margin-top: -0px;
}

.comparator{
  margin-top: 0px;
  overflow-y: auto;
  overflow-x: hidden;
  width: 90vw;
  margin-left: 5vw;
  margin-top: 10px;
  height: 77.3vh;
}

.calendar{
  position: absolute;
  margin-top: -73px;
  right: 22px;
  z-index: 500;
  border-radius: 10px;;
}

.sales{
  position: absolute;
  right: 200px;
  width: 100px;
  z-index: 200;
  margin-top: -15px;
}

.time{
  position: absolute;
  left: 300px;
}

</style>
