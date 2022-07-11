<template>
    <button :style="{background: color}" >{{text}}</button>
</template>

<script>
    export default{
        name: 'Button',
        props: {
            text: String,
            color: String
        },
        methods: {
        }
    }
</script>

<style scoped>
    button{
        padding: 10px 20px;
        text-align: center;
        text-decoration: none;
        font-size: 16px;
        border: none 2px;
        color: white;
        margin: 5px;
    }
</style>