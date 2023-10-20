<script setup>
import axios from "axios";
</script>

<template>
	<v-select id="opening-input" label="Opening" variant="outlined" v-model="hymnNumber" density="compact" :items="items" item-title="name" item-value="id" v-on:update:modelValue="onChange"></v-select>
</template>

<script>
export default {
	name: "Hymn",
	data() {
		return {
			hymnNumber: 0,
			items: [],
		}
	},
	mounted() {
		this.getHymns();
	},
	methods: {
		getHymns: function() {
			axios.get(import.meta.env.VITE_API_URL + "/v1/hymn?sort=id")
			.then(response => {
				this.items = response.data.data;
			})
			.catch(error => {
				console.log(error);
			});
		},
		onChange: function() {
			this.$emit("hymn-number-out", this.hymnNumber, this.items[this.hymnNumber-1])
		}
	}
}
</script>