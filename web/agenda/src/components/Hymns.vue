<script setup>
import Hymn from "./Hymn.vue";
</script>

<template>
	<v-row>
		<v-col cols="12" md="4">
			<Hymn :hymns="hymns" :hymn-number-in="opening" @hymn-number-out="openingHymn" label="Opening" />
		</v-col>
		<v-col cols="12" md="4">
			<Hymn :hymns="hymns" :hymn-number-in="sacrament" @hymn-number-out="sacramentHymn" label="Sacrament" />
		</v-col>
		<v-col cols="12" md="4">
			<Hymn :hymns="hymns" :hymn-number-in="closing" @hymn-number-out="closingHymn" label="Closing" />
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
			closing: 0,
			items: []
		}
	},
	props: [
		"agenda",
		"hymns"
	],
	emits: {
		refreshAgenda: null,
	},
	methods: {
		save: function() {
			const obj = {"date": this.agenda.date, "opening_hymn": this.opening, "sacrament_hymn": this.sacrament, "closing_hymn": this.closing};
			this.$emit("refreshAgenda", obj);
		},
		openingHymn(hymnNumber) {
			this.opening = hymnNumber
		},
		sacramentHymn(hymnNumber) {
			this.sacrament = hymnNumber
		},
		closingHymn(hymnNumber) {
			this.closing = hymnNumber
		}
	},
	watch: {
		agenda: {
		 	handler(newAgenda, oldAgenda) {
				this.opening = newAgenda.opening_hymn;
				this.sacrament = newAgenda.sacrament_hymn;
				this.closing = newAgenda.closing_hymn;
			},
			immediate: true
		}
	}
}
</script>