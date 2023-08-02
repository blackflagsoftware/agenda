<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="3">
			<v-text-field label="Name" variant="outlined" density="compact" v-model="localName"></v-text-field>
		</v-col>
		<v-col cols="12" md="5">
			<v-text-field label="Calling" variant="outlined" density="compact" v-model="localCalling"></v-text-field>
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
	name: "WardBusinessUpdate",
	data() {
		return {
			localName: this.item.name,
			localCalling: this.item.calling,
		}
	},
	props: [
		"item",
		"type"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "name": this.localName, "calling": this.localCalling}
			if (this.type === "release") {
				var url = import.meta.env.VITE_API_URL + "/v1/wardbusinessrel";
			} else {
				var url =  import.meta.env.VITE_API_URL + "/v1/wardbusinesssus";
			}
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshWardBusiness");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			if (this.type === "release") {
				var url = import.meta.env.VITE_API_URL + "/v1/wardbusinessrel/"+this.item.id
			} else {
				var url = import.meta.env.VITE_API_URL + "/v1/wardbusinesssus/"+this.item.id
			}
			axios.delete(url)
			.then(() => {
				this.$emit("refreshWardBusiness");
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.localName === "" || this.localCalling === "") ? true : false;
		}
	},
	watch: {
		item: {
			handler(newItem, oldItem) {
				this.localName = newItem.name;
				this.localCalling = newItem.calling;
			},
			immediate: true
		}
	}
}
</script>