<script setup>
import axios from "axios"
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="12">
				<v-text-field label="Confirmations" variant="outlined" density="compact" v-model="confirmation" @blur="confirmationBlur"></v-text-field>
			</v-col>
		</v-row>
		<v-row>
			<v-col cols="12" md="12">
				<v-text-field label="Blessings" variant="outlined" density="compact" v-model="blessing" @blur="blessingBlur"></v-text-field>
			</v-col>
		</v-row>
	</v-form>
</template>

<script>
export default {
	name: "Ordinance",
	data() {
		return {
			id: 0,
			confirmation: "",
			blessing: "",
		}
	},
	props: [
		"date"
	],
	mounted() {
		this.getOrdinance()
	},
	methods: {
		getOrdinance: function() {
			axios.get(import.meta.env.VITE_API_URL + "/v1/ordinance/"+this.date)
			.then(response => {
				this.id = response.data.data.id
				this.confirmation = response.data.data.confirmations
				this.blessing = response.data.data.blessings
			})
			.catch(error => {
				console.log(error)
			})
		},
		confirmationBlur: function() {
			const obj = {id: this.id, date: this.date, confirmations: this.confirmation}
			axios.post(import.meta.env.VITE_API_URL + "/v1/ordinance", obj)
			.then(response => {
				this.id = response.data.data.id
			})
			.catch(error => {
				console.log(error)
			})
		},
		blessingBlur: function() {
			const obj = {id: this.id, date: this.date, blessings: this.blessing}
			axios.post(import.meta.env.VITE_API_URL + "/v1/ordinance", obj)
			.then(response => {
				this.id = response.data.data.id
			})
			.catch(error => {
				console.log(error)
			})
		}
	}
}
</script>