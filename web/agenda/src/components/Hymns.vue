<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="4">
			<v-select id="opening-input" label="Opening" variant="outlined" v-model="opening" density="compact" :items="items" item-title="name" item-value="id"></v-select>
		</v-col>
		<v-col cols="12" md="4">
			<v-select id="sacrament-input" label="Sacrament" variant="outlined" v-model="sacrament" density="compact" :items="items" item-title="name" item-value="id"></v-select>
		</v-col>
		<v-col cols="12" md="4">
			<v-select id="intermediate-input" label="Intermediate" variant="outlined" v-model="intermediate" density="compact" :items="items" item-title="name" item-value="id"></v-select>
		</v-col>
		<v-col cols="12" md="4">
			<v-select id="closing-input" label="Closing" variant="outlined" v-model="closing" density="compact" item-title="name" :items="items" item-value="id"></v-select>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="12">
			<v-text-field id="musical-number-input" label="Musical Number" variant="outlined" v-model="musical_number" density="compact"></v-text-field>
		</v-col>
	</v-row>
	<v-row>
		<v-col>
			<v-btn @click="save" color="blue">Save</v-btn>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "Hymn",
	data() {
		return {
			opening: 0,
			sacrament: 0,
			intermediate: 0,
			closing: 0,
			musical_number: "",
			items: []
		}
	},
	mounted() {
		this.getHymns();
	},
	props: ["agenda"],
	emits: {
		refreshAgenda: null,
	},
	methods: {
		getHymns: function() {
			axios.get(import.meta.env.VITE_API_URL + "/v1/hymn?sort=id")
			.then(response => {
				this.items = response.data.data;
			})
			.catch(error => {
				console.log(error);
			});
		},
		save: function() {
			const obj = {"date": this.agenda.date, "opening_hymn": this.opening, "sacrament_hymn": this.sacrament, "intermediate_hymn": this.intermediate, "closing_hymn": this.closing, "musical_number": this.musical_number};
			this.$emit("refreshAgenda", obj);
		}
	},
	watch: {
		agenda: {
		 	handler(newAgenda, oldAgenda) {
				this.opening = newAgenda.opening_hymn;
				this.sacrament = newAgenda.sacrament_hymn;
				this.intermediate = newAgenda.intermediate_hymn;
				this.closing = newAgenda.closing_hymn;
				this.musical_number = newAgenda.musical_number;
			},
			immediate: true
		}
	}
}
</script>