<script setup>
import WardBusiness from "./WardBusiness.vue";
import BishopBusiness from "./BishopBusiness.vue";
import NewMember from "./NewMember.vue";
import Ordinance from "./Ordinance.vue";
</script>

<template>
	<v-row>
		<v-col cols="12" md="12">
			<v-checkbox-btn v-model="enabledWard" label="Ward" @change="wardChange"></v-checkbox-btn>
			<div style="padding-left:40px; padding-right:20px; padding-top:10px;">
				<WardBusiness v-if="enabledWard" :date="agenda.date"/>
			</div>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="12">
			<v-checkbox-btn v-model="enabledBishop" label="Bishop" @change="bishopChange"></v-checkbox-btn>
			<div style="padding-left:40px; padding-right:20px; padding-top:10px;">
				<BishopBusiness v-if="enabledBishop" :date="agenda.date"/>
			</div>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="12">
			<v-checkbox-btn v-model="enabledNewMember" label="New Members" @change="newMemberChange"></v-checkbox-btn>
			<div style="padding-left:40px; padding-right:20px; padding-top:10px;">
				<NewMember v-if="enabledNewMember" :date="agenda.date"/>
			</div>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="12">
			<v-checkbox-btn v-model="enabledOrdinance" label="Ordinances" @change="ordinanceChange"></v-checkbox-btn>
			<div style="padding-left:40px; padding-right:20px; padding-top:10px;">
				<Ordinance v-if="enabledOrdinance" :date="agenda.date"/>
			</div>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="12">
			<v-checkbox label="Stake Business" v-model="enabledStake" @change="stakeChange"></v-checkbox>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="12">
			<v-checkbox label="Read Letter" v-model="enabledLetter" @change="letterChange"></v-checkbox>
		</v-col>
	</v-row>
</template>

<script>
  export default {
	name: "Business",
    data: () => ({
      enabledWard: false,
	  enabledBishop: false,
	  enabledLetter: false,
	  enabledStake: false,
	  enabledNewMember: false,
	  enabledOrdinance: false,
    }),
	props: ["agenda"],
	emits: {
		refreshAgenda: null,
	},
	methods: {
		wardChange: function() {
			this.$emit("refreshAgenda", {"date": this.agenda.date, "ward_business": this.enabledWard});
		},
		bishopChange: function() {
			this.$emit("refreshAgenda", {"date": this.agenda.date, "bishop_business": this.enabledBishop});
		},
		newMemberChange: function() {
			this.$emit("refreshAgenda", {"date": this.agenda.date, "new_members": this.enabledNewMember});
		},
		ordinanceChange: function() {
			this.$emit("refreshAgenda", {"date": this.agenda.date, "ordinance": this.enabledOrdinance});
		},
		stakeChange: function() {
			this.$emit("refreshAgenda", {"date": this.agenda.date, "stake_business": this.enabledStake});
		},
		letterChange: function() {
			this.$emit("refreshAgenda", {"date": this.agenda.date, "letter_read": this.enabledLetter});
		}
	},
	watch: {
		agenda: {
		 	handler(newAgenda, oldAgenda) {
				this.enabledWard = newAgenda.ward_business;
				this.enabledBishop = newAgenda.bishop_business;
				this.enabledLetter = newAgenda.letter_read;
				this.enabledStake = newAgenda.stake_business;
				this.enabledNewMember = newAgenda.new_members;
				this.enabledOrdinance = newAgenda.ordinance;
			},
			immediate: true
		}
	}
}
</script>