<script setup>
import AnnouncementUpdate from "./AnnouncementUpdate.vue"
import axios from "axios"
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="6">
				<v-textarea auto-grow rows="1" label="Message" variant="outlined" density="compact" v-model="message"></v-textarea>
			</v-col>
			<v-col cols="12" md="4">
				<v-row align-self="center">
					<v-col cols="12" md="11">
						<v-text-field auto-grow rows="1" label="URL Link" variant="outlined" density="compact" v-model="url_link"></v-text-field>
					</v-col>
					<v-col cols="12" md="1" align-self="start">
						<v-tooltip location="bottom" text="Including a URL link will put the link after the annoucement on the QR version and a QR code on the printed program" >
        					<template v-slot:activator="{ props }">
								<v-btn v-bind="props" class="ma-2" variant="text" icon="mdi-information"></v-btn>
							</template>
						</v-tooltip>
					</v-col>
				</v-row>
			</v-col>
			<v-col cols="12" md="2">
				<v-row>
					<v-col cols="12" md="8">
						<v-checkbox label="Pulpit" v-model="pulpit"></v-checkbox>
					</v-col>
					<v-col cols="12" md="4">
						<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
					</v-col>
				</v-row>
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
			url_link: "",
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
			const obj = {"message": this.message, "url_link": this.url_link, "pulpit": this.pulpit}
			axios.post(import.meta.env.VITE_API_URL + "/v1/announcement", obj)
			.then(() => {
				this.message = "";
				this.url_link = "";
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