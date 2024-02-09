<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="6">
			<v-textarea auto-grow rows="1" label="Message" variant="outlined" density="compact" v-model="message"></v-textarea>
		</v-col>
		<v-col cols="12" md="3">
			<v-textarea auto-grow rows="1" label="URL Link" variant="outlined" density="compact" v-model="url_link"></v-textarea>
		</v-col>
		<v-col cols="12" md="3">
			<v-row>
				<v-col cols="12" md="4">
					<v-checkbox label="Pulpit" v-model="pulpit"></v-checkbox>
				</v-col>
				<v-col cols="12" md="4">
					<v-btn @click="updateSave()" :disabled="disableUpdateBtn" color="blue">Update</v-btn>
				</v-col>
				<v-col cols="12" md="4">
					<v-btn @click="updateDelete()" color="red">Delete</v-btn>
				</v-col>
			</v-row>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "AnnouncementUpdate",
	data() {
		return {
			pulpit: this.item.pulpit,
			message: this.item.message,
			url_link: this.item.url_link
		}
	},
	props: [
		"item"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "message": this.message, "url_link": this.url_link, "pulpit": this.pulpit};
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