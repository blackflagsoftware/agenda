package main

import (
	"net"
	"os"

	"github.com/blackflagsoftware/agenda/config"
	m "github.com/blackflagsoftware/agenda/internal/middleware"
	age "github.com/blackflagsoftware/agenda/internal/v1/agenda"
	ann "github.com/blackflagsoftware/agenda/internal/v1/announcement"
	bis "github.com/blackflagsoftware/agenda/internal/v1/bishopbusiness"
	def "github.com/blackflagsoftware/agenda/internal/v1/defaultcalling"
	hym "github.com/blackflagsoftware/agenda/internal/v1/hymn"
	mem "github.com/blackflagsoftware/agenda/internal/v1/member"
	new "github.com/blackflagsoftware/agenda/internal/v1/newmember"
	ord "github.com/blackflagsoftware/agenda/internal/v1/ordinance"
	rol "github.com/blackflagsoftware/agenda/internal/v1/role"
	ro "github.com/blackflagsoftware/agenda/internal/v1/roleuser"
	spe "github.com/blackflagsoftware/agenda/internal/v1/speaker"
	vis "github.com/blackflagsoftware/agenda/internal/v1/visitor"
	war "github.com/blackflagsoftware/agenda/internal/v1/wardbusinessrel"
	wa "github.com/blackflagsoftware/agenda/internal/v1/wardbusinesssus"
	pb "github.com/blackflagsoftware/agenda/pkg/proto"
	mig "github.com/blackflagsoftware/agenda/tools/migration/src"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// --- replace grpc import - do not remove ---
)

func main() {
	if config.UseMigration {
		err := os.MkdirAll(config.MigrationPath, 0744)
		if err != nil {
			m.Default.Printf("Unable to make scripts/migrations directory structure: %s\n", err)
		}
		c := mig.Connection{
			Host:           config.SqlitePath,
			MigrationPath:  config.MigrationPath,
			SkipInitialize: config.MigrationSkipInit,
			Engine:         "sqlite",
		}
		if err := mig.StartMigration(c); err != nil {
			m.Default.Panicf("Migration failed due to: %s", err)
		}
	}

	tcpListener, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		m.Default.Panic("Unable to start GRPC port:", err)
	}
	defer tcpListener.Close()
	s := grpc.NewServer()

	// Agenda
	sage := age.InitStorage()
	mage := age.NewManagerAgenda(sage)
	hage := age.NewAgendaGrpc(mage)
	pb.RegisterAgendaServiceServer(s, hage)
	// Visitor
	svis := vis.InitStorage()
	mvis := vis.NewManagerVisitor(svis)
	hvis := vis.NewVisitorGrpc(mvis)
	pb.RegisterVisitorServiceServer(s, hvis)
	// WardBusinessRel
	swar := war.InitStorage()
	mwar := war.NewManagerWardBusinessRel(swar)
	hwar := war.NewWardBusinessRelGrpc(mwar)
	pb.RegisterWardBusinessRelServiceServer(s, hwar)
	// WardBusinessSus
	swa := wa.InitStorage()
	mwa := wa.NewManagerWardBusinessSus(swa)
	hwa := wa.NewWardBusinessSusGrpc(mwa)
	pb.RegisterWardBusinessSusServiceServer(s, hwa)
	// BishopBusiness
	sbis := bis.InitStorage()
	mbis := bis.NewManagerBishopBusiness(sbis)
	hbis := bis.NewBishopBusinessGrpc(mbis)
	pb.RegisterBishopBusinessServiceServer(s, hbis)
	// NewMember
	snew := new.InitStorage()
	mnew := new.NewManagerNewMember(snew)
	hnew := new.NewNewMemberGrpc(mnew)
	pb.RegisterNewMemberServiceServer(s, hnew)
	// Ordinance
	sord := ord.InitStorage()
	mord := ord.NewManagerOrdinance(sord)
	hord := ord.NewOrdinanceGrpc(mord)
	pb.RegisterOrdinanceServiceServer(s, hord)
	// Speaker
	sspe := spe.InitStorage()
	mspe := spe.NewManagerSpeaker(sspe)
	hspe := spe.NewSpeakerGrpc(mspe)
	pb.RegisterSpeakerServiceServer(s, hspe)
	// Announcement
	sann := ann.InitStorage()
	mann := ann.NewManagerAnnouncement(sann)
	hann := ann.NewAnnouncementGrpc(mann)
	pb.RegisterAnnouncementServiceServer(s, hann)
	// Hymn
	shym := hym.InitStorage()
	mhym := hym.NewManagerHymn(shym)
	hhym := hym.NewHymnGrpc(mhym)
	pb.RegisterHymnServiceServer(s, hhym)
	// DefaultCalling
	sdef := def.InitStorage()
	mdef := def.NewManagerDefaultCalling(sdef)
	hdef := def.NewDefaultCallingGrpc(mdef)
	pb.RegisterDefaultCallingServiceServer(s, hdef)
	// Role
	srol := rol.InitStorage()
	mrol := rol.NewManagerRole(srol)
	hrol := rol.NewRoleGrpc(mrol)
	pb.RegisterRoleServiceServer(s, hrol)
	// RoleUser
	sro := ro.InitStorage()
	mro := ro.NewManagerRoleUser(sro)
	hro := ro.NewRoleUserGrpc(mro)
	pb.RegisterRoleUserServiceServer(s, hro)
	// Member
	smem := mem.InitStorage()
	mmem := mem.NewManagerMember(smem)
	hmem := mem.NewMemberGrpc(mmem)
	pb.RegisterMembersServiceServer(s, hmem)
	// --- replace grpc text - do not remove ---

	reflection.Register(s)
	m.Default.Printf("Starting GRPC server on port: %s...\n", config.GrpcPort)
	s.Serve(tcpListener)
}
