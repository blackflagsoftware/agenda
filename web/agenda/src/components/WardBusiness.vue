<script setup>
import WardBusinessUpdate from './WardBusinessUpdate.vue'
import axios from "axios"
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="2">
				<v-select label="Release/Sustain" :items="['Release', 'Sustaining']"  variant="outlined" density="compact" v-model="addWardType"></v-select>
			</v-col>
			<v-col cols="12" md="3">
				<v-text-field label="Name" variant="outlined" density="compact" v-model="addName"></v-text-field>
			</v-col>
			<v-col cols="12" md="6">
				<v-text-field label="Calling" variant="outlined" density="compact" v-model="addCalling"></v-text-field>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
			</v-col>
		</v-row>
	</v-form>
	<div style="margin-bottom:10px">
		<b>Releases</b>
	</div>
	<v-sheet v-if="showNoReleases">No Releases</v-sheet>
	<v-form>
		<WardBusinessUpdate v-for="r in releases" :type="'release'" :item="r" @refresh-ward-business="getWardBusiness"/>
	</v-form>
	<div style="margin-bottom:10px">
		<b>Sustaining</b>
	</div>
	<v-sheet v-if="showNoSustaining">No Sustainings</v-sheet>
	<v-form>
		<WardBusinessUpdate v-for="r in sustainings" :type="'sustaining'" :item="r" @refresh-ward-business="getWardBusiness"/>
	</v-form>
</template>

<script>
export default {
	name: "WardBusiness",
	data() {
		return {
			addWardType: "",
			addName: "",
			addCalling: "",
			releases: [],
			sustainings: [],
		}
	},
	props: [
		"date"
	],
	mounted() {
		this.getWardBusiness();
	},
	methods: {
		getWardBusiness: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/wardbusinessrel/search", {"search": [{"column": "date", "value": this.date, "compare": "="}]})
			.then(response => {
				this.releases = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
			axios.post(import.meta.env.VITE_API_URL + "/v1/wardbusinesssus/search", {"search": [{"column": "date", "value": this.date, "compare": "="}]})
			.then(response => {
				this.sustainings = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		addSave: function() {
			if (this.addWardType === "Release") {
				const obj = {"date": this.date, "name": this.addName, "calling": this.addCalling}
				axios.post(import.meta.env.VITE_API_URL + "/v1/wardbusinessrel", obj)
				.then(() => {
					this.addWardType = "";
					this.addName = "";
					this.addCalling = "";
					this.getWardBusiness();
				})
				.catch(error => {
					console.log(error);
				})
			} else {
				const obj = {"date": this.date, "name": this.addName, "calling": this.addCalling}
				axios.post(import.meta.env.VITE_API_URL + "/v1/wardbusinesssus", obj)
				.then(() => {
					this.getWardBusiness();
				})
				.catch(error => {
					console.log(error);
				})
			}
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.addWardType === "" || this.addName === "" || this.addCalling === "") ? true : false;
		},
		showNoReleases() {
			return this.releases.length === 0 ? true : false;
		},
		showNoSustaining() {
			return this.sustainings.length === 0 ? true : false;
		}
	}
}
</script>