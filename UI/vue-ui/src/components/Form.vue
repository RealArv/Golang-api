<template>
    <div class="form">
        <form @submit.prevent="airportRequest" method="POST">
        <div class="mb-3 row">
            <label for="" class="col-sm-2 col-form-label">Enter airport code</label>
            <div class="col-sm-10">
                <input type="text" v-model="formData.code" class="form-control">
            </div>
        </div>
        <div class="mb-3 row">
             <label for="" class="col-sm-2 col-form-label">Upload JSON file</label>
             <div class="col-sm-10">
                 <input type="file" @change="selectFile" ref="file">
             </div>
        </div>
            <Button @click="flush" type="submit" text="View results" color="green" class="m-3"/>
            <!-- <Button @click="flush" text="Clear results" color="blue"/> -->
        </form>

        <table class="table table-hover">
            <thead>
                <tr>
                    <th>Total Flights</th>
                    <th>Total Minutes Delayed</th>
                    <th>Time Label</th>
                </tr>
            </thead>
            <thead>
                <tr v-for="get in formData" :key="get.id">
                    <td>{{get.Total_Flights}}</td>
                    <td>{{get.Total_Minutes_Delayed}}</td>
                    <td>{{get.Time_Label}}</td>
                </tr>
            </thead>
        </table>
    </div>
</template>

<script>
    import axios from 'axios'
    import Button from './Button.vue'

    export default{
        name: 'Form',
        components: {
            Button,
        },

        data() {
            return {
                formData: [],    
            }
        },

        mounted() {
            // this.formatGetDetails();
            // this.getAirportInfo();
        },

        methods: {
            // async getAirportInfo() {
            //     const response = await axios.get('http://localhost:9090/get')
            //     // .then(response => {
            //         // })
            //     this.formatGetDetails(response.data)
            //         // console.log(response)
            // },
            formatGetDetails(gets) {
                for (let key in gets) {
                    this.formData.push({ ...gets[key], id:key})
                }
                // console.log(this.formData)
            },

            selectFile() {   
             this.formData.file = this.$refs.file.files[0]
            },

            airportRequest() {
                if(!this.formData.code){
                    alert('Please input airport code')
                    return
                }

                var demo = new FormData()
                demo.append('code', this.formData.code)
                demo.append('file', this.formData.file)

                axios
                    .post('http://localhost:9090/post', demo)
                    .then(response => {
                        let content = JSON.parse(JSON.stringify(response))
                        this.formatGetDetails(content.data.file)
                        // this.formData = content.data
                        // console.log(content.data)
                    })
                    .catch(error => console.log(error))
            },

            flush() {
                this.formData.splice(0)
            }
        }
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