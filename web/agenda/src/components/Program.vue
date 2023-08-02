<script setup>
import ProgramUpdate from "./ProgramUpdate.vue"
import axios from "axios"
</script>

<template>
	<div style="margin-bottom:10px">
		<b>Speakers</b>
	</div>
    <v-checkbox label="Is this a Fast Sunday?" v-model="fast_sunday" @change="clickFastSunday"></v-checkbox>
	<div v-if="!fast_sunday">
		<v-form>
			<v-row>
				<v-col cols="12" md="1">
					<v-select label="Position" variant="outlined" density="compact" v-model="position" :items="['1', '2', '3', '4', '5']"></v-select>
				</v-col>
				<v-col cols="12" md="10">
					<v-combobox label="Name" variant="outlined" density="compact" v-model="name" :items="['Intermediate Hymn', 'Musical Number']"></v-combobox>
				</v-col>
				<v-col cols="12" md="1">
					<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
				</v-col>
			</v-row>
		</v-form>
		<v-sheet v-if="showNoSpeakers">No Speakers</v-sheet>
		<v-form>
			<ProgramUpdate v-for="s in speakers" :item="s" @refresh-speakers="getSpeakers"/>
		</v-form>
	</div>
</template>

<script>
export default {
	name: "Program",
	data() {
		return {
            fast_sunday: false,
			localAgenda: null,
			position: "",
			name: "",
			speakers: []
		}
	},
	props: [
		"agenda"
	],
	mounted() {
		this.getSpeakers();
	},
	methods: {
		clickFastSunday: function() {
			axios.patch(import.meta.env.VITE_API_URL + "/v1/agenda", {"date": this.localAgenda.date, "fast_sunday": this.fast_sunday})
			.catch(error => {
				console.log(error);
			})
		},
		getSpeakers: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/speaker/search", {"search": [{"column": "date", "value": this.localAgenda.date, "compare": "="}]})
			.then(response => {
				this.speakers = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		addSave: function() {
			const obj = {"date": this.localAgenda.date, "position": this.position, "name": this.name}
			axios.post(import.meta.env.VITE_API_URL + "/v1/speaker", obj)
			.then(() => {
				this.position = "";
				this.name = "";
				this.getSpeakers();
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.position === "" || this.name === "") ? true : false;
		},
		showNoSpeakers() {
			return this.speakers.length === 0 ? true : false;
		}
	},
	watch: {
		agenda: {
			handler(newAgenda, oldAgenda) {
				this.localAgenda = newAgenda;
				this.getSpeakers();
			},
			immediate: true
		}
	}
}
</script>