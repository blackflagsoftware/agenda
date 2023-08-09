<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="8">
			<v-textarea auto-grow rows="1" label="Message" variant="outlined" density="compact" v-model="message"></v-textarea>
		</v-col>
		<v-col cols="12" md="2">
			<v-checkbox label="Pulpit" v-model="pulpit"></v-checkbox>
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
	name: "AnnouncementUpdate",
	data() {
		return {
			pulpit: this.item.pulpit,
			message: this.item.message
		}
	},
	props: [
		"item"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "message": this.message, "pulpit": this.pulpit};
			console.log("in announdement upate", obj);
			var url = import.meta.env.VITE_API_URL + "/v1/announcement";
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshAnnouncements");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			var url = import.meta.env.VITE_API_URL + "/v1/announcement/"+this.item.id
			axios.delete(url)
			.then(() => {
				this.$emit("refreshAnnouncements");
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.message === "") ? true : false;
		}
	}
}
</script>