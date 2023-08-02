<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="10">
			<v-text-field label="Name" variant="outlined" density="compact" v-model="localName"></v-text-field>
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
	name: "VisitorUpdate",
	data() {
		return {
			localName: this.item.name
		}
	},
	props: [
		"item"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "name": this.localName}
			var url = import.meta.env.VITE_API_URL + "/v1/visitor";
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshVisitor");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			var url = import.meta.env.VITE_API_URL + "/v1/visitor/"+this.item.id
			axios.delete(url)
			.then(() => {
				this.$emit("refreshVisitor");
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.localName === "") ? true : false;
		}
	},
	watch: {
		item: {
			handler(newItem, oldItem) {
				this.localName = newItem.name;
			},
			immediate: true
		}
	}
}
</script>