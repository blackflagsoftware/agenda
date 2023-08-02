<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="10">
			<v-text-field label="Message" variant="outlined" density="compact" v-model="localMessage"></v-text-field>
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
	name: "BishopBusinessUpdate",
	data() {
		return {
			localMessage: this.item.message
		}
	},
	props: [
		"item"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "message": this.localMessage}
			var url = import.meta.env.VITE_API_URL + "/v1/bishopbusiness";
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshBishop");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			var url = import.meta.env.VITE_API_URL + "/v1/bishopbusiness/"+this.item.id
			axios.delete(url)
			.then(() => {
				this.$emit("refreshBishop");
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.localMessage === "") ? true : false;
		}
	},
	watch: {
		item: {
			handler(newItem, oldItem) {
				this.localMessage = newItem.message;
			},
			immediate: true
		}
	}
}
</script>