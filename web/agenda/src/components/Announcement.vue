<script setup>
import AnnouncementUpdate from "./AnnouncementUpdate.vue"
import axios from "axios"
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="9">
				<v-textarea auto-grow rows="1" label="Message" variant="outlined" density="compact" v-model="message"></v-textarea>
			</v-col>
			<v-col cols="12" md="2">
				<v-checkbox label="Pulpit" v-model="pulpit"></v-checkbox>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
			</v-col>
		</v-row>
	</v-form>
	<v-sheet v-if="showNoAnnouncements">No Announcements</v-sheet>
	<v-form>
		<AnnouncementUpdate v-for="a in announcements" v-bind:item="a" v-bind:key="a.id" @refresh-announcements="getAnnouncements"/>
	</v-form>
</template>

<script>
export default {
	name: "Announcement",
	data() {
		return {
			pulpit: false,
			message: "",
			announcements: []
		}
	},
	mounted() {
		this.getAnnouncements();
	},
	methods: {
		getAnnouncements: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/announcement/search", {"search": []})
			.then(response => {
				this.announcements = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		addSave: function() {
			const obj = {"message": this.message, "pulpit": this.pulpit}
			axios.post(import.meta.env.VITE_API_URL + "/v1/announcement", obj)
			.then(() => {
				this.message = "";
				this.getAnnouncements();
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.message === "") ? true : false;
		},
		showNoAnnouncements() {
			return this.announcements.length === 0 ? true : false;
		}
	}
}
</script>