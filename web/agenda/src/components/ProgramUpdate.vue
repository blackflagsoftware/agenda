<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="1">
			<v-select label="Position" variant="outlined" density="compact" v-model="position" :items="['1', '2', '3', '4', '5']"></v-select>
		</v-col>
		<v-col cols="12" md="9">
			<v-combobox label="Name" variant="outlined" density="compact" v-model="name" :items="['Intermediate Hymn', 'Musical Number']"></v-combobox>
		</v-col>
		<v-col cols="12" md="1">
			<v-btn @click="updateSave()" :disabled="disableUpdateBtn" color="blue">Update</v-btn>
		</v-col>
		<v-col cols="12" md="1">
			<v-btn @click="updateDelete()" color="red">Delete</v-btn>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "ProgramUpdate",
	data() {
		return {
			position: this.item.position,
			name: this.item.name
		}
	},
	props: [
		"item"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "position": this.position, "name": this.name}
			var url = import.meta.env.VITE_API_URL + "/v1/speaker";
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshSpeakers");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			var url = import.meta.env.VITE_API_URL + "/v1/speaker/"+this.item.id
			axios.delete(url)
			.then(() => {
				this.$emit("refreshSpeakers");
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.position === "" || this.name === "") ? true : false;
		}
	},
	watch: {
		item: {
		 	handler(newItem, oldItem) {
				this.position = newItem.position;
				this.name = newItem.name;
			},
			immediate: true
		}
	}
}
</script>