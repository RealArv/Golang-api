<template>
    <div class="form">
        <form @submit="CreatePost" method="POST">
            <label for="">Enter airport code</label>
            <input type="text" v-model="posts.code">
            <br> <br>
            <label for="">Upload JSON file</label>
            <input type="file">
            <br><br>
            <Button type="submit" text="View results" color="green"/>
            <!-- <Button @click="getClick" text="get button" color="blue"/> -->
        </form>

        <table>
            <tr>
                <th>Total Flights</th>
                <th>Total Minutes Delayed</th>
                <th>Time Label</th>
            </tr>
            <tr v-for="get in getDetails" :key="get.id">
                <td>{{get.Total_Flights}}</td>
                <td>{{get.Total_Minutes_Delayed}}</td>
                <td>{{get.Time_Label}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    import axios from 'axios'
    import { ref } from 'vue'
    
    import Button from './Button.vue'

    export default{
        name: 'Form',
        components: {
            Button,
        },

        data() {
            return {
                getDetails: [],
                posts: {
                    code: ''
                },
            }
        },

        mounted() {
            this.formatGetDetails();
            this.getAirportInfo();
        },

        methods: {
            onCreatePost() {
                axios.post(
                    'http://localhost:9090/post',
                        {code: this.code, file: this.file}
                    ).then(response => {
                        console.log(response);
                    })
            },

            async getAirportInfo() {
                const response = await axios.get('http://localhost:9090/get')
                // .then(response => {
                    // })
                this.formatGetDetails(response.data)
                    // console.log(response)
            },
            formatGetDetails(gets) {
                for (let key in gets) {
                    this.getDetails.push({ ...gets[key], id:key})
                }
                console.log(this.getDetails)
            }

    }

        // data() {
        //     return {
        //         contents: null,
        //     }
        // },
        // mounted() {
        //     axios
        //         .get('http://localhost:9090/get')
        //         .then((response) => (this.contents = response.data))
        //         .catch(error => {
        //             this.errorMsg = error.message
        //             console.log("Error", error)
        //         })
        // }
        // setup() {
        //     const res = ref('')

        //     axios.get('http://localhost:9090/get')
        //     .then(response => {
        //         res.value = response.data.message
        //         res.value2 = response.data.id
        //         console.log(response)
        //     })

        //     return {
        //         res
        //     }
        // }
    }
</script>

<style scoped>
    .form{
        margin: 30px;
        font-size: 20px;
    }
    input{
        margin-left: 10px;
    }
</style>