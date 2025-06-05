<template>
    <v-card class="dialog1">
        <p class="question">{{ question }}</p>
        <div class="buttons">
            <v-btn v-for="(caption, n) in buttons" :key="n" :index="n" @click="close(caption)">{{ caption }}</v-btn>
        </div>
    </v-card>

</template>

<script>
    import {VBtn, VCard} from "vuetify/lib";

    export default{
        name: "ConfirmDialog",
        components:{
            VBtn,
            VCard,
        },
        props:{
            question: String,
            buttons: Array,
        },
        methods:{
            close(code){
                this.$emit("close", code);
            }
        }
    }
</script>

<style scoped>
    .dialog1{
        position: absolute;
        width: 500px;
        height: 150px;
        border-style: solid;
        border-width: 5px;
        border-color: #ff5e00;
        border-radius: 25px;
        background: #90a0bf;
        left: 50%;
        margin-left: -250px;
        top: 50%;
        margin-top: -75px;
        padding: 30px;
    }

    .question{
        color:white;
    }

    .buttons{
        width: 100%;
        margin-left: 0px;
        text-align: center;
    }

    .v-btn{
        text-align: center;
        margin-left: 10px;
        border-style: solid;
        border-color: #ff5e00;
        border-radius: 15px;
    }

</style>