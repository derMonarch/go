// DO NOT EDIT
// This file is auto-generated by the following steps, and the sysreggen.go is located at x/arch/arm64/arm64spec/.
// 1. Get the system register xml files from https://developer.arm.com/-/media/Files/ATG/Beta10/SysReg_xml_v85A-2019-06.tar.gz
// 2. Extract SysReg_xml_v85A-2019-06.tar/SysReg_xml_v85A-2019-06/SysReg_xml_v85A-2019-06/AArch64-*.xml to ./files folder
// 3. Run the command: go run sysreggen.go

package arm64

const (
	AUTO_SYSREG_BEGIN = REG_SPECIAL + iota
	REG_ACTLR_EL1
	REG_AFSR0_EL1
	REG_AFSR1_EL1
	REG_AIDR_EL1
	REG_AMAIR_EL1
	REG_AMCFGR_EL0
	REG_AMCGCR_EL0
	REG_AMCNTENCLR0_EL0
	REG_AMCNTENCLR1_EL0
	REG_AMCNTENSET0_EL0
	REG_AMCNTENSET1_EL0
	REG_AMCR_EL0
	REG_AMEVCNTR00_EL0
	REG_AMEVCNTR01_EL0
	REG_AMEVCNTR02_EL0
	REG_AMEVCNTR03_EL0
	REG_AMEVCNTR04_EL0
	REG_AMEVCNTR05_EL0
	REG_AMEVCNTR06_EL0
	REG_AMEVCNTR07_EL0
	REG_AMEVCNTR08_EL0
	REG_AMEVCNTR09_EL0
	REG_AMEVCNTR010_EL0
	REG_AMEVCNTR011_EL0
	REG_AMEVCNTR012_EL0
	REG_AMEVCNTR013_EL0
	REG_AMEVCNTR014_EL0
	REG_AMEVCNTR015_EL0
	REG_AMEVCNTR10_EL0
	REG_AMEVCNTR11_EL0
	REG_AMEVCNTR12_EL0
	REG_AMEVCNTR13_EL0
	REG_AMEVCNTR14_EL0
	REG_AMEVCNTR15_EL0
	REG_AMEVCNTR16_EL0
	REG_AMEVCNTR17_EL0
	REG_AMEVCNTR18_EL0
	REG_AMEVCNTR19_EL0
	REG_AMEVCNTR110_EL0
	REG_AMEVCNTR111_EL0
	REG_AMEVCNTR112_EL0
	REG_AMEVCNTR113_EL0
	REG_AMEVCNTR114_EL0
	REG_AMEVCNTR115_EL0
	REG_AMEVTYPER00_EL0
	REG_AMEVTYPER01_EL0
	REG_AMEVTYPER02_EL0
	REG_AMEVTYPER03_EL0
	REG_AMEVTYPER04_EL0
	REG_AMEVTYPER05_EL0
	REG_AMEVTYPER06_EL0
	REG_AMEVTYPER07_EL0
	REG_AMEVTYPER08_EL0
	REG_AMEVTYPER09_EL0
	REG_AMEVTYPER010_EL0
	REG_AMEVTYPER011_EL0
	REG_AMEVTYPER012_EL0
	REG_AMEVTYPER013_EL0
	REG_AMEVTYPER014_EL0
	REG_AMEVTYPER015_EL0
	REG_AMEVTYPER10_EL0
	REG_AMEVTYPER11_EL0
	REG_AMEVTYPER12_EL0
	REG_AMEVTYPER13_EL0
	REG_AMEVTYPER14_EL0
	REG_AMEVTYPER15_EL0
	REG_AMEVTYPER16_EL0
	REG_AMEVTYPER17_EL0
	REG_AMEVTYPER18_EL0
	REG_AMEVTYPER19_EL0
	REG_AMEVTYPER110_EL0
	REG_AMEVTYPER111_EL0
	REG_AMEVTYPER112_EL0
	REG_AMEVTYPER113_EL0
	REG_AMEVTYPER114_EL0
	REG_AMEVTYPER115_EL0
	REG_AMUSERENR_EL0
	REG_APDAKeyHi_EL1
	REG_APDAKeyLo_EL1
	REG_APDBKeyHi_EL1
	REG_APDBKeyLo_EL1
	REG_APGAKeyHi_EL1
	REG_APGAKeyLo_EL1
	REG_APIAKeyHi_EL1
	REG_APIAKeyLo_EL1
	REG_APIBKeyHi_EL1
	REG_APIBKeyLo_EL1
	REG_CCSIDR2_EL1
	REG_CCSIDR_EL1
	REG_CLIDR_EL1
	REG_CNTFRQ_EL0
	REG_CNTKCTL_EL1
	REG_CNTP_CTL_EL0
	REG_CNTP_CVAL_EL0
	REG_CNTP_TVAL_EL0
	REG_CNTPCT_EL0
	REG_CNTPS_CTL_EL1
	REG_CNTPS_CVAL_EL1
	REG_CNTPS_TVAL_EL1
	REG_CNTV_CTL_EL0
	REG_CNTV_CVAL_EL0
	REG_CNTV_TVAL_EL0
	REG_CNTVCT_EL0
	REG_CONTEXTIDR_EL1
	REG_CPACR_EL1
	REG_CSSELR_EL1
	REG_CTR_EL0
	REG_CurrentEL
	REG_DAIF
	REG_DBGAUTHSTATUS_EL1
	REG_DBGBCR0_EL1
	REG_DBGBCR1_EL1
	REG_DBGBCR2_EL1
	REG_DBGBCR3_EL1
	REG_DBGBCR4_EL1
	REG_DBGBCR5_EL1
	REG_DBGBCR6_EL1
	REG_DBGBCR7_EL1
	REG_DBGBCR8_EL1
	REG_DBGBCR9_EL1
	REG_DBGBCR10_EL1
	REG_DBGBCR11_EL1
	REG_DBGBCR12_EL1
	REG_DBGBCR13_EL1
	REG_DBGBCR14_EL1
	REG_DBGBCR15_EL1
	REG_DBGBVR0_EL1
	REG_DBGBVR1_EL1
	REG_DBGBVR2_EL1
	REG_DBGBVR3_EL1
	REG_DBGBVR4_EL1
	REG_DBGBVR5_EL1
	REG_DBGBVR6_EL1
	REG_DBGBVR7_EL1
	REG_DBGBVR8_EL1
	REG_DBGBVR9_EL1
	REG_DBGBVR10_EL1
	REG_DBGBVR11_EL1
	REG_DBGBVR12_EL1
	REG_DBGBVR13_EL1
	REG_DBGBVR14_EL1
	REG_DBGBVR15_EL1
	REG_DBGCLAIMCLR_EL1
	REG_DBGCLAIMSET_EL1
	REG_DBGDTR_EL0
	REG_DBGDTRRX_EL0
	REG_DBGDTRTX_EL0
	REG_DBGPRCR_EL1
	REG_DBGWCR0_EL1
	REG_DBGWCR1_EL1
	REG_DBGWCR2_EL1
	REG_DBGWCR3_EL1
	REG_DBGWCR4_EL1
	REG_DBGWCR5_EL1
	REG_DBGWCR6_EL1
	REG_DBGWCR7_EL1
	REG_DBGWCR8_EL1
	REG_DBGWCR9_EL1
	REG_DBGWCR10_EL1
	REG_DBGWCR11_EL1
	REG_DBGWCR12_EL1
	REG_DBGWCR13_EL1
	REG_DBGWCR14_EL1
	REG_DBGWCR15_EL1
	REG_DBGWVR0_EL1
	REG_DBGWVR1_EL1
	REG_DBGWVR2_EL1
	REG_DBGWVR3_EL1
	REG_DBGWVR4_EL1
	REG_DBGWVR5_EL1
	REG_DBGWVR6_EL1
	REG_DBGWVR7_EL1
	REG_DBGWVR8_EL1
	REG_DBGWVR9_EL1
	REG_DBGWVR10_EL1
	REG_DBGWVR11_EL1
	REG_DBGWVR12_EL1
	REG_DBGWVR13_EL1
	REG_DBGWVR14_EL1
	REG_DBGWVR15_EL1
	REG_DCZID_EL0
	REG_DISR_EL1
	REG_DIT
	REG_DLR_EL0
	REG_DSPSR_EL0
	REG_ELR_EL1
	REG_ERRIDR_EL1
	REG_ERRSELR_EL1
	REG_ERXADDR_EL1
	REG_ERXCTLR_EL1
	REG_ERXFR_EL1
	REG_ERXMISC0_EL1
	REG_ERXMISC1_EL1
	REG_ERXMISC2_EL1
	REG_ERXMISC3_EL1
	REG_ERXPFGCDN_EL1
	REG_ERXPFGCTL_EL1
	REG_ERXPFGF_EL1
	REG_ERXSTATUS_EL1
	REG_ESR_EL1
	REG_FAR_EL1
	REG_FPCR
	REG_FPSR
	REG_GCR_EL1
	REG_GMID_EL1
	REG_ICC_AP0R0_EL1
	REG_ICC_AP0R1_EL1
	REG_ICC_AP0R2_EL1
	REG_ICC_AP0R3_EL1
	REG_ICC_AP1R0_EL1
	REG_ICC_AP1R1_EL1
	REG_ICC_AP1R2_EL1
	REG_ICC_AP1R3_EL1
	REG_ICC_ASGI1R_EL1
	REG_ICC_BPR0_EL1
	REG_ICC_BPR1_EL1
	REG_ICC_CTLR_EL1
	REG_ICC_DIR_EL1
	REG_ICC_EOIR0_EL1
	REG_ICC_EOIR1_EL1
	REG_ICC_HPPIR0_EL1
	REG_ICC_HPPIR1_EL1
	REG_ICC_IAR0_EL1
	REG_ICC_IAR1_EL1
	REG_ICC_IGRPEN0_EL1
	REG_ICC_IGRPEN1_EL1
	REG_ICC_PMR_EL1
	REG_ICC_RPR_EL1
	REG_ICC_SGI0R_EL1
	REG_ICC_SGI1R_EL1
	REG_ICC_SRE_EL1
	REG_ICV_AP0R0_EL1
	REG_ICV_AP0R1_EL1
	REG_ICV_AP0R2_EL1
	REG_ICV_AP0R3_EL1
	REG_ICV_AP1R0_EL1
	REG_ICV_AP1R1_EL1
	REG_ICV_AP1R2_EL1
	REG_ICV_AP1R3_EL1
	REG_ICV_BPR0_EL1
	REG_ICV_BPR1_EL1
	REG_ICV_CTLR_EL1
	REG_ICV_DIR_EL1
	REG_ICV_EOIR0_EL1
	REG_ICV_EOIR1_EL1
	REG_ICV_HPPIR0_EL1
	REG_ICV_HPPIR1_EL1
	REG_ICV_IAR0_EL1
	REG_ICV_IAR1_EL1
	REG_ICV_IGRPEN0_EL1
	REG_ICV_IGRPEN1_EL1
	REG_ICV_PMR_EL1
	REG_ICV_RPR_EL1
	REG_ID_AA64AFR0_EL1
	REG_ID_AA64AFR1_EL1
	REG_ID_AA64DFR0_EL1
	REG_ID_AA64DFR1_EL1
	REG_ID_AA64ISAR0_EL1
	REG_ID_AA64ISAR1_EL1
	REG_ID_AA64MMFR0_EL1
	REG_ID_AA64MMFR1_EL1
	REG_ID_AA64MMFR2_EL1
	REG_ID_AA64PFR0_EL1
	REG_ID_AA64PFR1_EL1
	REG_ID_AA64ZFR0_EL1
	REG_ID_AFR0_EL1
	REG_ID_DFR0_EL1
	REG_ID_ISAR0_EL1
	REG_ID_ISAR1_EL1
	REG_ID_ISAR2_EL1
	REG_ID_ISAR3_EL1
	REG_ID_ISAR4_EL1
	REG_ID_ISAR5_EL1
	REG_ID_ISAR6_EL1
	REG_ID_MMFR0_EL1
	REG_ID_MMFR1_EL1
	REG_ID_MMFR2_EL1
	REG_ID_MMFR3_EL1
	REG_ID_MMFR4_EL1
	REG_ID_PFR0_EL1
	REG_ID_PFR1_EL1
	REG_ID_PFR2_EL1
	REG_ISR_EL1
	REG_LORC_EL1
	REG_LOREA_EL1
	REG_LORID_EL1
	REG_LORN_EL1
	REG_LORSA_EL1
	REG_MAIR_EL1
	REG_MDCCINT_EL1
	REG_MDCCSR_EL0
	REG_MDRAR_EL1
	REG_MDSCR_EL1
	REG_MIDR_EL1
	REG_MPAM0_EL1
	REG_MPAM1_EL1
	REG_MPAMIDR_EL1
	REG_MPIDR_EL1
	REG_MVFR0_EL1
	REG_MVFR1_EL1
	REG_MVFR2_EL1
	REG_NZCV
	REG_OSDLR_EL1
	REG_OSDTRRX_EL1
	REG_OSDTRTX_EL1
	REG_OSECCR_EL1
	REG_OSLAR_EL1
	REG_OSLSR_EL1
	REG_PAN
	REG_PAR_EL1
	REG_PMBIDR_EL1
	REG_PMBLIMITR_EL1
	REG_PMBPTR_EL1
	REG_PMBSR_EL1
	REG_PMCCFILTR_EL0
	REG_PMCCNTR_EL0
	REG_PMCEID0_EL0
	REG_PMCEID1_EL0
	REG_PMCNTENCLR_EL0
	REG_PMCNTENSET_EL0
	REG_PMCR_EL0
	REG_PMEVCNTR0_EL0
	REG_PMEVCNTR1_EL0
	REG_PMEVCNTR2_EL0
	REG_PMEVCNTR3_EL0
	REG_PMEVCNTR4_EL0
	REG_PMEVCNTR5_EL0
	REG_PMEVCNTR6_EL0
	REG_PMEVCNTR7_EL0
	REG_PMEVCNTR8_EL0
	REG_PMEVCNTR9_EL0
	REG_PMEVCNTR10_EL0
	REG_PMEVCNTR11_EL0
	REG_PMEVCNTR12_EL0
	REG_PMEVCNTR13_EL0
	REG_PMEVCNTR14_EL0
	REG_PMEVCNTR15_EL0
	REG_PMEVCNTR16_EL0
	REG_PMEVCNTR17_EL0
	REG_PMEVCNTR18_EL0
	REG_PMEVCNTR19_EL0
	REG_PMEVCNTR20_EL0
	REG_PMEVCNTR21_EL0
	REG_PMEVCNTR22_EL0
	REG_PMEVCNTR23_EL0
	REG_PMEVCNTR24_EL0
	REG_PMEVCNTR25_EL0
	REG_PMEVCNTR26_EL0
	REG_PMEVCNTR27_EL0
	REG_PMEVCNTR28_EL0
	REG_PMEVCNTR29_EL0
	REG_PMEVCNTR30_EL0
	REG_PMEVTYPER0_EL0
	REG_PMEVTYPER1_EL0
	REG_PMEVTYPER2_EL0
	REG_PMEVTYPER3_EL0
	REG_PMEVTYPER4_EL0
	REG_PMEVTYPER5_EL0
	REG_PMEVTYPER6_EL0
	REG_PMEVTYPER7_EL0
	REG_PMEVTYPER8_EL0
	REG_PMEVTYPER9_EL0
	REG_PMEVTYPER10_EL0
	REG_PMEVTYPER11_EL0
	REG_PMEVTYPER12_EL0
	REG_PMEVTYPER13_EL0
	REG_PMEVTYPER14_EL0
	REG_PMEVTYPER15_EL0
	REG_PMEVTYPER16_EL0
	REG_PMEVTYPER17_EL0
	REG_PMEVTYPER18_EL0
	REG_PMEVTYPER19_EL0
	REG_PMEVTYPER20_EL0
	REG_PMEVTYPER21_EL0
	REG_PMEVTYPER22_EL0
	REG_PMEVTYPER23_EL0
	REG_PMEVTYPER24_EL0
	REG_PMEVTYPER25_EL0
	REG_PMEVTYPER26_EL0
	REG_PMEVTYPER27_EL0
	REG_PMEVTYPER28_EL0
	REG_PMEVTYPER29_EL0
	REG_PMEVTYPER30_EL0
	REG_PMINTENCLR_EL1
	REG_PMINTENSET_EL1
	REG_PMMIR_EL1
	REG_PMOVSCLR_EL0
	REG_PMOVSSET_EL0
	REG_PMSCR_EL1
	REG_PMSELR_EL0
	REG_PMSEVFR_EL1
	REG_PMSFCR_EL1
	REG_PMSICR_EL1
	REG_PMSIDR_EL1
	REG_PMSIRR_EL1
	REG_PMSLATFR_EL1
	REG_PMSWINC_EL0
	REG_PMUSERENR_EL0
	REG_PMXEVCNTR_EL0
	REG_PMXEVTYPER_EL0
	REG_REVIDR_EL1
	REG_RGSR_EL1
	REG_RMR_EL1
	REG_RNDR
	REG_RNDRRS
	REG_RVBAR_EL1
	REG_SCTLR_EL1
	REG_SCXTNUM_EL0
	REG_SCXTNUM_EL1
	REG_SP_EL0
	REG_SP_EL1
	REG_SPSel
	REG_SPSR_abt
	REG_SPSR_EL1
	REG_SPSR_fiq
	REG_SPSR_irq
	REG_SPSR_und
	REG_SSBS
	REG_TCO
	REG_TCR_EL1
	REG_TFSR_EL1
	REG_TFSRE0_EL1
	REG_TPIDR_EL0
	REG_TPIDR_EL1
	REG_TPIDRRO_EL0
	REG_TRFCR_EL1
	REG_TTBR0_EL1
	REG_TTBR1_EL1
	REG_UAO
	REG_VBAR_EL1
	REG_ZCR_EL1
	AUTO_SYSREG_END
)

var SystemReg = []struct {
	Name string
	Reg  int16
	Enc  uint32
}{
	{"ACTLR_EL1", REG_ACTLR_EL1, 0x181020},
	{"AFSR0_EL1", REG_AFSR0_EL1, 0x185100},
	{"AFSR1_EL1", REG_AFSR1_EL1, 0x185120},
	{"AIDR_EL1", REG_AIDR_EL1, 0x1900e0},
	{"AMAIR_EL1", REG_AMAIR_EL1, 0x18a300},
	{"AMCFGR_EL0", REG_AMCFGR_EL0, 0x1bd220},
	{"AMCGCR_EL0", REG_AMCGCR_EL0, 0x1bd240},
	{"AMCNTENCLR0_EL0", REG_AMCNTENCLR0_EL0, 0x1bd280},
	{"AMCNTENCLR1_EL0", REG_AMCNTENCLR1_EL0, 0x1bd300},
	{"AMCNTENSET0_EL0", REG_AMCNTENSET0_EL0, 0x1bd2a0},
	{"AMCNTENSET1_EL0", REG_AMCNTENSET1_EL0, 0x1bd320},
	{"AMCR_EL0", REG_AMCR_EL0, 0x1bd200},
	{"AMEVCNTR00_EL0", REG_AMEVCNTR00_EL0, 0x1bd400},
	{"AMEVCNTR01_EL0", REG_AMEVCNTR01_EL0, 0x1bd420},
	{"AMEVCNTR02_EL0", REG_AMEVCNTR02_EL0, 0x1bd440},
	{"AMEVCNTR03_EL0", REG_AMEVCNTR03_EL0, 0x1bd460},
	{"AMEVCNTR04_EL0", REG_AMEVCNTR04_EL0, 0x1bd480},
	{"AMEVCNTR05_EL0", REG_AMEVCNTR05_EL0, 0x1bd4a0},
	{"AMEVCNTR06_EL0", REG_AMEVCNTR06_EL0, 0x1bd4c0},
	{"AMEVCNTR07_EL0", REG_AMEVCNTR07_EL0, 0x1bd4e0},
	{"AMEVCNTR08_EL0", REG_AMEVCNTR08_EL0, 0x1bd500},
	{"AMEVCNTR09_EL0", REG_AMEVCNTR09_EL0, 0x1bd520},
	{"AMEVCNTR010_EL0", REG_AMEVCNTR010_EL0, 0x1bd540},
	{"AMEVCNTR011_EL0", REG_AMEVCNTR011_EL0, 0x1bd560},
	{"AMEVCNTR012_EL0", REG_AMEVCNTR012_EL0, 0x1bd580},
	{"AMEVCNTR013_EL0", REG_AMEVCNTR013_EL0, 0x1bd5a0},
	{"AMEVCNTR014_EL0", REG_AMEVCNTR014_EL0, 0x1bd5c0},
	{"AMEVCNTR015_EL0", REG_AMEVCNTR015_EL0, 0x1bd5e0},
	{"AMEVCNTR10_EL0", REG_AMEVCNTR10_EL0, 0x1bdc00},
	{"AMEVCNTR11_EL0", REG_AMEVCNTR11_EL0, 0x1bdc20},
	{"AMEVCNTR12_EL0", REG_AMEVCNTR12_EL0, 0x1bdc40},
	{"AMEVCNTR13_EL0", REG_AMEVCNTR13_EL0, 0x1bdc60},
	{"AMEVCNTR14_EL0", REG_AMEVCNTR14_EL0, 0x1bdc80},
	{"AMEVCNTR15_EL0", REG_AMEVCNTR15_EL0, 0x1bdca0},
	{"AMEVCNTR16_EL0", REG_AMEVCNTR16_EL0, 0x1bdcc0},
	{"AMEVCNTR17_EL0", REG_AMEVCNTR17_EL0, 0x1bdce0},
	{"AMEVCNTR18_EL0", REG_AMEVCNTR18_EL0, 0x1bdd00},
	{"AMEVCNTR19_EL0", REG_AMEVCNTR19_EL0, 0x1bdd20},
	{"AMEVCNTR110_EL0", REG_AMEVCNTR110_EL0, 0x1bdd40},
	{"AMEVCNTR111_EL0", REG_AMEVCNTR111_EL0, 0x1bdd60},
	{"AMEVCNTR112_EL0", REG_AMEVCNTR112_EL0, 0x1bdd80},
	{"AMEVCNTR113_EL0", REG_AMEVCNTR113_EL0, 0x1bdda0},
	{"AMEVCNTR114_EL0", REG_AMEVCNTR114_EL0, 0x1bddc0},
	{"AMEVCNTR115_EL0", REG_AMEVCNTR115_EL0, 0x1bdde0},
	{"AMEVTYPER00_EL0", REG_AMEVTYPER00_EL0, 0x1bd600},
	{"AMEVTYPER01_EL0", REG_AMEVTYPER01_EL0, 0x1bd620},
	{"AMEVTYPER02_EL0", REG_AMEVTYPER02_EL0, 0x1bd640},
	{"AMEVTYPER03_EL0", REG_AMEVTYPER03_EL0, 0x1bd660},
	{"AMEVTYPER04_EL0", REG_AMEVTYPER04_EL0, 0x1bd680},
	{"AMEVTYPER05_EL0", REG_AMEVTYPER05_EL0, 0x1bd6a0},
	{"AMEVTYPER06_EL0", REG_AMEVTYPER06_EL0, 0x1bd6c0},
	{"AMEVTYPER07_EL0", REG_AMEVTYPER07_EL0, 0x1bd6e0},
	{"AMEVTYPER08_EL0", REG_AMEVTYPER08_EL0, 0x1bd700},
	{"AMEVTYPER09_EL0", REG_AMEVTYPER09_EL0, 0x1bd720},
	{"AMEVTYPER010_EL0", REG_AMEVTYPER010_EL0, 0x1bd740},
	{"AMEVTYPER011_EL0", REG_AMEVTYPER011_EL0, 0x1bd760},
	{"AMEVTYPER012_EL0", REG_AMEVTYPER012_EL0, 0x1bd780},
	{"AMEVTYPER013_EL0", REG_AMEVTYPER013_EL0, 0x1bd7a0},
	{"AMEVTYPER014_EL0", REG_AMEVTYPER014_EL0, 0x1bd7c0},
	{"AMEVTYPER015_EL0", REG_AMEVTYPER015_EL0, 0x1bd7e0},
	{"AMEVTYPER10_EL0", REG_AMEVTYPER10_EL0, 0x1bde00},
	{"AMEVTYPER11_EL0", REG_AMEVTYPER11_EL0, 0x1bde20},
	{"AMEVTYPER12_EL0", REG_AMEVTYPER12_EL0, 0x1bde40},
	{"AMEVTYPER13_EL0", REG_AMEVTYPER13_EL0, 0x1bde60},
	{"AMEVTYPER14_EL0", REG_AMEVTYPER14_EL0, 0x1bde80},
	{"AMEVTYPER15_EL0", REG_AMEVTYPER15_EL0, 0x1bdea0},
	{"AMEVTYPER16_EL0", REG_AMEVTYPER16_EL0, 0x1bdec0},
	{"AMEVTYPER17_EL0", REG_AMEVTYPER17_EL0, 0x1bdee0},
	{"AMEVTYPER18_EL0", REG_AMEVTYPER18_EL0, 0x1bdf00},
	{"AMEVTYPER19_EL0", REG_AMEVTYPER19_EL0, 0x1bdf20},
	{"AMEVTYPER110_EL0", REG_AMEVTYPER110_EL0, 0x1bdf40},
	{"AMEVTYPER111_EL0", REG_AMEVTYPER111_EL0, 0x1bdf60},
	{"AMEVTYPER112_EL0", REG_AMEVTYPER112_EL0, 0x1bdf80},
	{"AMEVTYPER113_EL0", REG_AMEVTYPER113_EL0, 0x1bdfa0},
	{"AMEVTYPER114_EL0", REG_AMEVTYPER114_EL0, 0x1bdfc0},
	{"AMEVTYPER115_EL0", REG_AMEVTYPER115_EL0, 0x1bdfe0},
	{"AMUSERENR_EL0", REG_AMUSERENR_EL0, 0x1bd260},
	{"APDAKeyHi_EL1", REG_APDAKeyHi_EL1, 0x182220},
	{"APDAKeyLo_EL1", REG_APDAKeyLo_EL1, 0x182200},
	{"APDBKeyHi_EL1", REG_APDBKeyHi_EL1, 0x182260},
	{"APDBKeyLo_EL1", REG_APDBKeyLo_EL1, 0x182240},
	{"APGAKeyHi_EL1", REG_APGAKeyHi_EL1, 0x182320},
	{"APGAKeyLo_EL1", REG_APGAKeyLo_EL1, 0x182300},
	{"APIAKeyHi_EL1", REG_APIAKeyHi_EL1, 0x182120},
	{"APIAKeyLo_EL1", REG_APIAKeyLo_EL1, 0x182100},
	{"APIBKeyHi_EL1", REG_APIBKeyHi_EL1, 0x182160},
	{"APIBKeyLo_EL1", REG_APIBKeyLo_EL1, 0x182140},
	{"CCSIDR2_EL1", REG_CCSIDR2_EL1, 0x190040},
	{"CCSIDR_EL1", REG_CCSIDR_EL1, 0x190000},
	{"CLIDR_EL1", REG_CLIDR_EL1, 0x190020},
	{"CNTFRQ_EL0", REG_CNTFRQ_EL0, 0x1be000},
	{"CNTKCTL_EL1", REG_CNTKCTL_EL1, 0x18e100},
	{"CNTP_CTL_EL0", REG_CNTP_CTL_EL0, 0x1be220},
	{"CNTP_CVAL_EL0", REG_CNTP_CVAL_EL0, 0x1be240},
	{"CNTP_TVAL_EL0", REG_CNTP_TVAL_EL0, 0x1be200},
	{"CNTPCT_EL0", REG_CNTPCT_EL0, 0x1be020},
	{"CNTPS_CTL_EL1", REG_CNTPS_CTL_EL1, 0x1fe220},
	{"CNTPS_CVAL_EL1", REG_CNTPS_CVAL_EL1, 0x1fe240},
	{"CNTPS_TVAL_EL1", REG_CNTPS_TVAL_EL1, 0x1fe200},
	{"CNTV_CTL_EL0", REG_CNTV_CTL_EL0, 0x1be320},
	{"CNTV_CVAL_EL0", REG_CNTV_CVAL_EL0, 0x1be340},
	{"CNTV_TVAL_EL0", REG_CNTV_TVAL_EL0, 0x1be300},
	{"CNTVCT_EL0", REG_CNTVCT_EL0, 0x1be040},
	{"CONTEXTIDR_EL1", REG_CONTEXTIDR_EL1, 0x18d020},
	{"CPACR_EL1", REG_CPACR_EL1, 0x181040},
	{"CSSELR_EL1", REG_CSSELR_EL1, 0x1a0000},
	{"CTR_EL0", REG_CTR_EL0, 0x1b0020},
	{"CurrentEL", REG_CurrentEL, 0x184240},
	{"DAIF", REG_DAIF, 0x1b4220},
	{"DBGAUTHSTATUS_EL1", REG_DBGAUTHSTATUS_EL1, 0x107ec0},
	{"DBGBCR0_EL1", REG_DBGBCR0_EL1, 0x1000a0},
	{"DBGBCR1_EL1", REG_DBGBCR1_EL1, 0x1001a0},
	{"DBGBCR2_EL1", REG_DBGBCR2_EL1, 0x1002a0},
	{"DBGBCR3_EL1", REG_DBGBCR3_EL1, 0x1003a0},
	{"DBGBCR4_EL1", REG_DBGBCR4_EL1, 0x1004a0},
	{"DBGBCR5_EL1", REG_DBGBCR5_EL1, 0x1005a0},
	{"DBGBCR6_EL1", REG_DBGBCR6_EL1, 0x1006a0},
	{"DBGBCR7_EL1", REG_DBGBCR7_EL1, 0x1007a0},
	{"DBGBCR8_EL1", REG_DBGBCR8_EL1, 0x1008a0},
	{"DBGBCR9_EL1", REG_DBGBCR9_EL1, 0x1009a0},
	{"DBGBCR10_EL1", REG_DBGBCR10_EL1, 0x100aa0},
	{"DBGBCR11_EL1", REG_DBGBCR11_EL1, 0x100ba0},
	{"DBGBCR12_EL1", REG_DBGBCR12_EL1, 0x100ca0},
	{"DBGBCR13_EL1", REG_DBGBCR13_EL1, 0x100da0},
	{"DBGBCR14_EL1", REG_DBGBCR14_EL1, 0x100ea0},
	{"DBGBCR15_EL1", REG_DBGBCR15_EL1, 0x100fa0},
	{"DBGBVR0_EL1", REG_DBGBVR0_EL1, 0x100080},
	{"DBGBVR1_EL1", REG_DBGBVR1_EL1, 0x100180},
	{"DBGBVR2_EL1", REG_DBGBVR2_EL1, 0x100280},
	{"DBGBVR3_EL1", REG_DBGBVR3_EL1, 0x100380},
	{"DBGBVR4_EL1", REG_DBGBVR4_EL1, 0x100480},
	{"DBGBVR5_EL1", REG_DBGBVR5_EL1, 0x100580},
	{"DBGBVR6_EL1", REG_DBGBVR6_EL1, 0x100680},
	{"DBGBVR7_EL1", REG_DBGBVR7_EL1, 0x100780},
	{"DBGBVR8_EL1", REG_DBGBVR8_EL1, 0x100880},
	{"DBGBVR9_EL1", REG_DBGBVR9_EL1, 0x100980},
	{"DBGBVR10_EL1", REG_DBGBVR10_EL1, 0x100a80},
	{"DBGBVR11_EL1", REG_DBGBVR11_EL1, 0x100b80},
	{"DBGBVR12_EL1", REG_DBGBVR12_EL1, 0x100c80},
	{"DBGBVR13_EL1", REG_DBGBVR13_EL1, 0x100d80},
	{"DBGBVR14_EL1", REG_DBGBVR14_EL1, 0x100e80},
	{"DBGBVR15_EL1", REG_DBGBVR15_EL1, 0x100f80},
	{"DBGCLAIMCLR_EL1", REG_DBGCLAIMCLR_EL1, 0x1079c0},
	{"DBGCLAIMSET_EL1", REG_DBGCLAIMSET_EL1, 0x1078c0},
	{"DBGDTR_EL0", REG_DBGDTR_EL0, 0x130400},
	{"DBGDTRRX_EL0", REG_DBGDTRRX_EL0, 0x130500},
	{"DBGDTRTX_EL0", REG_DBGDTRTX_EL0, 0x130500},
	{"DBGPRCR_EL1", REG_DBGPRCR_EL1, 0x101480},
	{"DBGWCR0_EL1", REG_DBGWCR0_EL1, 0x1000e0},
	{"DBGWCR1_EL1", REG_DBGWCR1_EL1, 0x1001e0},
	{"DBGWCR2_EL1", REG_DBGWCR2_EL1, 0x1002e0},
	{"DBGWCR3_EL1", REG_DBGWCR3_EL1, 0x1003e0},
	{"DBGWCR4_EL1", REG_DBGWCR4_EL1, 0x1004e0},
	{"DBGWCR5_EL1", REG_DBGWCR5_EL1, 0x1005e0},
	{"DBGWCR6_EL1", REG_DBGWCR6_EL1, 0x1006e0},
	{"DBGWCR7_EL1", REG_DBGWCR7_EL1, 0x1007e0},
	{"DBGWCR8_EL1", REG_DBGWCR8_EL1, 0x1008e0},
	{"DBGWCR9_EL1", REG_DBGWCR9_EL1, 0x1009e0},
	{"DBGWCR10_EL1", REG_DBGWCR10_EL1, 0x100ae0},
	{"DBGWCR11_EL1", REG_DBGWCR11_EL1, 0x100be0},
	{"DBGWCR12_EL1", REG_DBGWCR12_EL1, 0x100ce0},
	{"DBGWCR13_EL1", REG_DBGWCR13_EL1, 0x100de0},
	{"DBGWCR14_EL1", REG_DBGWCR14_EL1, 0x100ee0},
	{"DBGWCR15_EL1", REG_DBGWCR15_EL1, 0x100fe0},
	{"DBGWVR0_EL1", REG_DBGWVR0_EL1, 0x1000c0},
	{"DBGWVR1_EL1", REG_DBGWVR1_EL1, 0x1001c0},
	{"DBGWVR2_EL1", REG_DBGWVR2_EL1, 0x1002c0},
	{"DBGWVR3_EL1", REG_DBGWVR3_EL1, 0x1003c0},
	{"DBGWVR4_EL1", REG_DBGWVR4_EL1, 0x1004c0},
	{"DBGWVR5_EL1", REG_DBGWVR5_EL1, 0x1005c0},
	{"DBGWVR6_EL1", REG_DBGWVR6_EL1, 0x1006c0},
	{"DBGWVR7_EL1", REG_DBGWVR7_EL1, 0x1007c0},
	{"DBGWVR8_EL1", REG_DBGWVR8_EL1, 0x1008c0},
	{"DBGWVR9_EL1", REG_DBGWVR9_EL1, 0x1009c0},
	{"DBGWVR10_EL1", REG_DBGWVR10_EL1, 0x100ac0},
	{"DBGWVR11_EL1", REG_DBGWVR11_EL1, 0x100bc0},
	{"DBGWVR12_EL1", REG_DBGWVR12_EL1, 0x100cc0},
	{"DBGWVR13_EL1", REG_DBGWVR13_EL1, 0x100dc0},
	{"DBGWVR14_EL1", REG_DBGWVR14_EL1, 0x100ec0},
	{"DBGWVR15_EL1", REG_DBGWVR15_EL1, 0x100fc0},
	{"DCZID_EL0", REG_DCZID_EL0, 0x1b00e0},
	{"DISR_EL1", REG_DISR_EL1, 0x18c120},
	{"DIT", REG_DIT, 0x1b42a0},
	{"DLR_EL0", REG_DLR_EL0, 0x1b4520},
	{"DSPSR_EL0", REG_DSPSR_EL0, 0x1b4500},
	{"ELR_EL1", REG_ELR_EL1, 0x184020},
	{"ERRIDR_EL1", REG_ERRIDR_EL1, 0x185300},
	{"ERRSELR_EL1", REG_ERRSELR_EL1, 0x185320},
	{"ERXADDR_EL1", REG_ERXADDR_EL1, 0x185460},
	{"ERXCTLR_EL1", REG_ERXCTLR_EL1, 0x185420},
	{"ERXFR_EL1", REG_ERXFR_EL1, 0x185400},
	{"ERXMISC0_EL1", REG_ERXMISC0_EL1, 0x185500},
	{"ERXMISC1_EL1", REG_ERXMISC1_EL1, 0x185520},
	{"ERXMISC2_EL1", REG_ERXMISC2_EL1, 0x185540},
	{"ERXMISC3_EL1", REG_ERXMISC3_EL1, 0x185560},
	{"ERXPFGCDN_EL1", REG_ERXPFGCDN_EL1, 0x1854c0},
	{"ERXPFGCTL_EL1", REG_ERXPFGCTL_EL1, 0x1854a0},
	{"ERXPFGF_EL1", REG_ERXPFGF_EL1, 0x185480},
	{"ERXSTATUS_EL1", REG_ERXSTATUS_EL1, 0x185440},
	{"ESR_EL1", REG_ESR_EL1, 0x185200},
	{"FAR_EL1", REG_FAR_EL1, 0x186000},
	{"FPCR", REG_FPCR, 0x1b4400},
	{"FPSR", REG_FPSR, 0x1b4420},
	{"GCR_EL1", REG_GCR_EL1, 0x1810c0},
	{"GMID_EL1", REG_GMID_EL1, 0x31400},
	{"ICC_AP0R0_EL1", REG_ICC_AP0R0_EL1, 0x18c880},
	{"ICC_AP0R1_EL1", REG_ICC_AP0R1_EL1, 0x18c8a0},
	{"ICC_AP0R2_EL1", REG_ICC_AP0R2_EL1, 0x18c8c0},
	{"ICC_AP0R3_EL1", REG_ICC_AP0R3_EL1, 0x18c8e0},
	{"ICC_AP1R0_EL1", REG_ICC_AP1R0_EL1, 0x18c900},
	{"ICC_AP1R1_EL1", REG_ICC_AP1R1_EL1, 0x18c920},
	{"ICC_AP1R2_EL1", REG_ICC_AP1R2_EL1, 0x18c940},
	{"ICC_AP1R3_EL1", REG_ICC_AP1R3_EL1, 0x18c960},
	{"ICC_ASGI1R_EL1", REG_ICC_ASGI1R_EL1, 0x18cbc0},
	{"ICC_BPR0_EL1", REG_ICC_BPR0_EL1, 0x18c860},
	{"ICC_BPR1_EL1", REG_ICC_BPR1_EL1, 0x18cc60},
	{"ICC_CTLR_EL1", REG_ICC_CTLR_EL1, 0x18cc80},
	{"ICC_DIR_EL1", REG_ICC_DIR_EL1, 0x18cb20},
	{"ICC_EOIR0_EL1", REG_ICC_EOIR0_EL1, 0x18c820},
	{"ICC_EOIR1_EL1", REG_ICC_EOIR1_EL1, 0x18cc20},
	{"ICC_HPPIR0_EL1", REG_ICC_HPPIR0_EL1, 0x18c840},
	{"ICC_HPPIR1_EL1", REG_ICC_HPPIR1_EL1, 0x18cc40},
	{"ICC_IAR0_EL1", REG_ICC_IAR0_EL1, 0x18c800},
	{"ICC_IAR1_EL1", REG_ICC_IAR1_EL1, 0x18cc00},
	{"ICC_IGRPEN0_EL1", REG_ICC_IGRPEN0_EL1, 0x18ccc0},
	{"ICC_IGRPEN1_EL1", REG_ICC_IGRPEN1_EL1, 0x18cce0},
	{"ICC_PMR_EL1", REG_ICC_PMR_EL1, 0x184600},
	{"ICC_RPR_EL1", REG_ICC_RPR_EL1, 0x18cb60},
	{"ICC_SGI0R_EL1", REG_ICC_SGI0R_EL1, 0x18cbe0},
	{"ICC_SGI1R_EL1", REG_ICC_SGI1R_EL1, 0x18cba0},
	{"ICC_SRE_EL1", REG_ICC_SRE_EL1, 0x18cca0},
	{"ICV_AP0R0_EL1", REG_ICV_AP0R0_EL1, 0x18c880},
	{"ICV_AP0R1_EL1", REG_ICV_AP0R1_EL1, 0x18c8a0},
	{"ICV_AP0R2_EL1", REG_ICV_AP0R2_EL1, 0x18c8c0},
	{"ICV_AP0R3_EL1", REG_ICV_AP0R3_EL1, 0x18c8e0},
	{"ICV_AP1R0_EL1", REG_ICV_AP1R0_EL1, 0x18c900},
	{"ICV_AP1R1_EL1", REG_ICV_AP1R1_EL1, 0x18c920},
	{"ICV_AP1R2_EL1", REG_ICV_AP1R2_EL1, 0x18c940},
	{"ICV_AP1R3_EL1", REG_ICV_AP1R3_EL1, 0x18c960},
	{"ICV_BPR0_EL1", REG_ICV_BPR0_EL1, 0x18c860},
	{"ICV_BPR1_EL1", REG_ICV_BPR1_EL1, 0x18cc60},
	{"ICV_CTLR_EL1", REG_ICV_CTLR_EL1, 0x18cc80},
	{"ICV_DIR_EL1", REG_ICV_DIR_EL1, 0x18cb20},
	{"ICV_EOIR0_EL1", REG_ICV_EOIR0_EL1, 0x18c820},
	{"ICV_EOIR1_EL1", REG_ICV_EOIR1_EL1, 0x18cc20},
	{"ICV_HPPIR0_EL1", REG_ICV_HPPIR0_EL1, 0x18c840},
	{"ICV_HPPIR1_EL1", REG_ICV_HPPIR1_EL1, 0x18cc40},
	{"ICV_IAR0_EL1", REG_ICV_IAR0_EL1, 0x18c800},
	{"ICV_IAR1_EL1", REG_ICV_IAR1_EL1, 0x18cc00},
	{"ICV_IGRPEN0_EL1", REG_ICV_IGRPEN0_EL1, 0x18ccc0},
	{"ICV_IGRPEN1_EL1", REG_ICV_IGRPEN1_EL1, 0x18cce0},
	{"ICV_PMR_EL1", REG_ICV_PMR_EL1, 0x184600},
	{"ICV_RPR_EL1", REG_ICV_RPR_EL1, 0x18cb60},
	{"ID_AA64AFR0_EL1", REG_ID_AA64AFR0_EL1, 0x180580},
	{"ID_AA64AFR1_EL1", REG_ID_AA64AFR1_EL1, 0x1805a0},
	{"ID_AA64DFR0_EL1", REG_ID_AA64DFR0_EL1, 0x180500},
	{"ID_AA64DFR1_EL1", REG_ID_AA64DFR1_EL1, 0x180520},
	{"ID_AA64ISAR0_EL1", REG_ID_AA64ISAR0_EL1, 0x180600},
	{"ID_AA64ISAR1_EL1", REG_ID_AA64ISAR1_EL1, 0x180620},
	{"ID_AA64MMFR0_EL1", REG_ID_AA64MMFR0_EL1, 0x180700},
	{"ID_AA64MMFR1_EL1", REG_ID_AA64MMFR1_EL1, 0x180720},
	{"ID_AA64MMFR2_EL1", REG_ID_AA64MMFR2_EL1, 0x180740},
	{"ID_AA64PFR0_EL1", REG_ID_AA64PFR0_EL1, 0x180400},
	{"ID_AA64PFR1_EL1", REG_ID_AA64PFR1_EL1, 0x180420},
	{"ID_AA64ZFR0_EL1", REG_ID_AA64ZFR0_EL1, 0x180480},
	{"ID_AFR0_EL1", REG_ID_AFR0_EL1, 0x180160},
	{"ID_DFR0_EL1", REG_ID_DFR0_EL1, 0x180140},
	{"ID_ISAR0_EL1", REG_ID_ISAR0_EL1, 0x180200},
	{"ID_ISAR1_EL1", REG_ID_ISAR1_EL1, 0x180220},
	{"ID_ISAR2_EL1", REG_ID_ISAR2_EL1, 0x180240},
	{"ID_ISAR3_EL1", REG_ID_ISAR3_EL1, 0x180260},
	{"ID_ISAR4_EL1", REG_ID_ISAR4_EL1, 0x180280},
	{"ID_ISAR5_EL1", REG_ID_ISAR5_EL1, 0x1802a0},
	{"ID_ISAR6_EL1", REG_ID_ISAR6_EL1, 0x1802e0},
	{"ID_MMFR0_EL1", REG_ID_MMFR0_EL1, 0x180180},
	{"ID_MMFR1_EL1", REG_ID_MMFR1_EL1, 0x1801a0},
	{"ID_MMFR2_EL1", REG_ID_MMFR2_EL1, 0x1801c0},
	{"ID_MMFR3_EL1", REG_ID_MMFR3_EL1, 0x1801e0},
	{"ID_MMFR4_EL1", REG_ID_MMFR4_EL1, 0x1802c0},
	{"ID_PFR0_EL1", REG_ID_PFR0_EL1, 0x180100},
	{"ID_PFR1_EL1", REG_ID_PFR1_EL1, 0x180120},
	{"ID_PFR2_EL1", REG_ID_PFR2_EL1, 0x180380},
	{"ISR_EL1", REG_ISR_EL1, 0x18c100},
	{"LORC_EL1", REG_LORC_EL1, 0x18a460},
	{"LOREA_EL1", REG_LOREA_EL1, 0x18a420},
	{"LORID_EL1", REG_LORID_EL1, 0x18a4e0},
	{"LORN_EL1", REG_LORN_EL1, 0x18a440},
	{"LORSA_EL1", REG_LORSA_EL1, 0x18a400},
	{"MAIR_EL1", REG_MAIR_EL1, 0x18a200},
	{"MDCCINT_EL1", REG_MDCCINT_EL1, 0x100200},
	{"MDCCSR_EL0", REG_MDCCSR_EL0, 0x130100},
	{"MDRAR_EL1", REG_MDRAR_EL1, 0x101000},
	{"MDSCR_EL1", REG_MDSCR_EL1, 0x100240},
	{"MIDR_EL1", REG_MIDR_EL1, 0x180000},
	{"MPAM0_EL1", REG_MPAM0_EL1, 0x18a520},
	{"MPAM1_EL1", REG_MPAM1_EL1, 0x18a500},
	{"MPAMIDR_EL1", REG_MPAMIDR_EL1, 0x18a480},
	{"MPIDR_EL1", REG_MPIDR_EL1, 0x1800a0},
	{"MVFR0_EL1", REG_MVFR0_EL1, 0x180300},
	{"MVFR1_EL1", REG_MVFR1_EL1, 0x180320},
	{"MVFR2_EL1", REG_MVFR2_EL1, 0x180340},
	{"NZCV", REG_NZCV, 0x1b4200},
	{"OSDLR_EL1", REG_OSDLR_EL1, 0x101380},
	{"OSDTRRX_EL1", REG_OSDTRRX_EL1, 0x100040},
	{"OSDTRTX_EL1", REG_OSDTRTX_EL1, 0x100340},
	{"OSECCR_EL1", REG_OSECCR_EL1, 0x100640},
	{"OSLAR_EL1", REG_OSLAR_EL1, 0x101080},
	{"OSLSR_EL1", REG_OSLSR_EL1, 0x101180},
	{"PAN", REG_PAN, 0x184260},
	{"PAR_EL1", REG_PAR_EL1, 0x187400},
	{"PMBIDR_EL1", REG_PMBIDR_EL1, 0x189ae0},
	{"PMBLIMITR_EL1", REG_PMBLIMITR_EL1, 0x189a00},
	{"PMBPTR_EL1", REG_PMBPTR_EL1, 0x189a20},
	{"PMBSR_EL1", REG_PMBSR_EL1, 0x189a60},
	{"PMCCFILTR_EL0", REG_PMCCFILTR_EL0, 0x1befe0},
	{"PMCCNTR_EL0", REG_PMCCNTR_EL0, 0x1b9d00},
	{"PMCEID0_EL0", REG_PMCEID0_EL0, 0x1b9cc0},
	{"PMCEID1_EL0", REG_PMCEID1_EL0, 0x1b9ce0},
	{"PMCNTENCLR_EL0", REG_PMCNTENCLR_EL0, 0x1b9c40},
	{"PMCNTENSET_EL0", REG_PMCNTENSET_EL0, 0x1b9c20},
	{"PMCR_EL0", REG_PMCR_EL0, 0x1b9c00},
	{"PMEVCNTR0_EL0", REG_PMEVCNTR0_EL0, 0x1be800},
	{"PMEVCNTR1_EL0", REG_PMEVCNTR1_EL0, 0x1be820},
	{"PMEVCNTR2_EL0", REG_PMEVCNTR2_EL0, 0x1be840},
	{"PMEVCNTR3_EL0", REG_PMEVCNTR3_EL0, 0x1be860},
	{"PMEVCNTR4_EL0", REG_PMEVCNTR4_EL0, 0x1be880},
	{"PMEVCNTR5_EL0", REG_PMEVCNTR5_EL0, 0x1be8a0},
	{"PMEVCNTR6_EL0", REG_PMEVCNTR6_EL0, 0x1be8c0},
	{"PMEVCNTR7_EL0", REG_PMEVCNTR7_EL0, 0x1be8e0},
	{"PMEVCNTR8_EL0", REG_PMEVCNTR8_EL0, 0x1be900},
	{"PMEVCNTR9_EL0", REG_PMEVCNTR9_EL0, 0x1be920},
	{"PMEVCNTR10_EL0", REG_PMEVCNTR10_EL0, 0x1be940},
	{"PMEVCNTR11_EL0", REG_PMEVCNTR11_EL0, 0x1be960},
	{"PMEVCNTR12_EL0", REG_PMEVCNTR12_EL0, 0x1be980},
	{"PMEVCNTR13_EL0", REG_PMEVCNTR13_EL0, 0x1be9a0},
	{"PMEVCNTR14_EL0", REG_PMEVCNTR14_EL0, 0x1be9c0},
	{"PMEVCNTR15_EL0", REG_PMEVCNTR15_EL0, 0x1be9e0},
	{"PMEVCNTR16_EL0", REG_PMEVCNTR16_EL0, 0x1bea00},
	{"PMEVCNTR17_EL0", REG_PMEVCNTR17_EL0, 0x1bea20},
	{"PMEVCNTR18_EL0", REG_PMEVCNTR18_EL0, 0x1bea40},
	{"PMEVCNTR19_EL0", REG_PMEVCNTR19_EL0, 0x1bea60},
	{"PMEVCNTR20_EL0", REG_PMEVCNTR20_EL0, 0x1bea80},
	{"PMEVCNTR21_EL0", REG_PMEVCNTR21_EL0, 0x1beaa0},
	{"PMEVCNTR22_EL0", REG_PMEVCNTR22_EL0, 0x1beac0},
	{"PMEVCNTR23_EL0", REG_PMEVCNTR23_EL0, 0x1beae0},
	{"PMEVCNTR24_EL0", REG_PMEVCNTR24_EL0, 0x1beb00},
	{"PMEVCNTR25_EL0", REG_PMEVCNTR25_EL0, 0x1beb20},
	{"PMEVCNTR26_EL0", REG_PMEVCNTR26_EL0, 0x1beb40},
	{"PMEVCNTR27_EL0", REG_PMEVCNTR27_EL0, 0x1beb60},
	{"PMEVCNTR28_EL0", REG_PMEVCNTR28_EL0, 0x1beb80},
	{"PMEVCNTR29_EL0", REG_PMEVCNTR29_EL0, 0x1beba0},
	{"PMEVCNTR30_EL0", REG_PMEVCNTR30_EL0, 0x1bebc0},
	{"PMEVTYPER0_EL0", REG_PMEVTYPER0_EL0, 0x1bec00},
	{"PMEVTYPER1_EL0", REG_PMEVTYPER1_EL0, 0x1bec20},
	{"PMEVTYPER2_EL0", REG_PMEVTYPER2_EL0, 0x1bec40},
	{"PMEVTYPER3_EL0", REG_PMEVTYPER3_EL0, 0x1bec60},
	{"PMEVTYPER4_EL0", REG_PMEVTYPER4_EL0, 0x1bec80},
	{"PMEVTYPER5_EL0", REG_PMEVTYPER5_EL0, 0x1beca0},
	{"PMEVTYPER6_EL0", REG_PMEVTYPER6_EL0, 0x1becc0},
	{"PMEVTYPER7_EL0", REG_PMEVTYPER7_EL0, 0x1bece0},
	{"PMEVTYPER8_EL0", REG_PMEVTYPER8_EL0, 0x1bed00},
	{"PMEVTYPER9_EL0", REG_PMEVTYPER9_EL0, 0x1bed20},
	{"PMEVTYPER10_EL0", REG_PMEVTYPER10_EL0, 0x1bed40},
	{"PMEVTYPER11_EL0", REG_PMEVTYPER11_EL0, 0x1bed60},
	{"PMEVTYPER12_EL0", REG_PMEVTYPER12_EL0, 0x1bed80},
	{"PMEVTYPER13_EL0", REG_PMEVTYPER13_EL0, 0x1beda0},
	{"PMEVTYPER14_EL0", REG_PMEVTYPER14_EL0, 0x1bedc0},
	{"PMEVTYPER15_EL0", REG_PMEVTYPER15_EL0, 0x1bede0},
	{"PMEVTYPER16_EL0", REG_PMEVTYPER16_EL0, 0x1bee00},
	{"PMEVTYPER17_EL0", REG_PMEVTYPER17_EL0, 0x1bee20},
	{"PMEVTYPER18_EL0", REG_PMEVTYPER18_EL0, 0x1bee40},
	{"PMEVTYPER19_EL0", REG_PMEVTYPER19_EL0, 0x1bee60},
	{"PMEVTYPER20_EL0", REG_PMEVTYPER20_EL0, 0x1bee80},
	{"PMEVTYPER21_EL0", REG_PMEVTYPER21_EL0, 0x1beea0},
	{"PMEVTYPER22_EL0", REG_PMEVTYPER22_EL0, 0x1beec0},
	{"PMEVTYPER23_EL0", REG_PMEVTYPER23_EL0, 0x1beee0},
	{"PMEVTYPER24_EL0", REG_PMEVTYPER24_EL0, 0x1bef00},
	{"PMEVTYPER25_EL0", REG_PMEVTYPER25_EL0, 0x1bef20},
	{"PMEVTYPER26_EL0", REG_PMEVTYPER26_EL0, 0x1bef40},
	{"PMEVTYPER27_EL0", REG_PMEVTYPER27_EL0, 0x1bef60},
	{"PMEVTYPER28_EL0", REG_PMEVTYPER28_EL0, 0x1bef80},
	{"PMEVTYPER29_EL0", REG_PMEVTYPER29_EL0, 0x1befa0},
	{"PMEVTYPER30_EL0", REG_PMEVTYPER30_EL0, 0x1befc0},
	{"PMINTENCLR_EL1", REG_PMINTENCLR_EL1, 0x189e40},
	{"PMINTENSET_EL1", REG_PMINTENSET_EL1, 0x189e20},
	{"PMMIR_EL1", REG_PMMIR_EL1, 0x189ec0},
	{"PMOVSCLR_EL0", REG_PMOVSCLR_EL0, 0x1b9c60},
	{"PMOVSSET_EL0", REG_PMOVSSET_EL0, 0x1b9e60},
	{"PMSCR_EL1", REG_PMSCR_EL1, 0x189900},
	{"PMSELR_EL0", REG_PMSELR_EL0, 0x1b9ca0},
	{"PMSEVFR_EL1", REG_PMSEVFR_EL1, 0x1899a0},
	{"PMSFCR_EL1", REG_PMSFCR_EL1, 0x189980},
	{"PMSICR_EL1", REG_PMSICR_EL1, 0x189940},
	{"PMSIDR_EL1", REG_PMSIDR_EL1, 0x1899e0},
	{"PMSIRR_EL1", REG_PMSIRR_EL1, 0x189960},
	{"PMSLATFR_EL1", REG_PMSLATFR_EL1, 0x1899c0},
	{"PMSWINC_EL0", REG_PMSWINC_EL0, 0x1b9c80},
	{"PMUSERENR_EL0", REG_PMUSERENR_EL0, 0x1b9e00},
	{"PMXEVCNTR_EL0", REG_PMXEVCNTR_EL0, 0x1b9d40},
	{"PMXEVTYPER_EL0", REG_PMXEVTYPER_EL0, 0x1b9d20},
	{"REVIDR_EL1", REG_REVIDR_EL1, 0x1800c0},
	{"RGSR_EL1", REG_RGSR_EL1, 0x1810a0},
	{"RMR_EL1", REG_RMR_EL1, 0x18c040},
	{"RNDR", REG_RNDR, 0x1b2400},
	{"RNDRRS", REG_RNDRRS, 0x1b2420},
	{"RVBAR_EL1", REG_RVBAR_EL1, 0x18c020},
	{"SCTLR_EL1", REG_SCTLR_EL1, 0x181000},
	{"SCXTNUM_EL0", REG_SCXTNUM_EL0, 0x1bd0e0},
	{"SCXTNUM_EL1", REG_SCXTNUM_EL1, 0x18d0e0},
	{"SP_EL0", REG_SP_EL0, 0x184100},
	{"SP_EL1", REG_SP_EL1, 0x1c4100},
	{"SPSel", REG_SPSel, 0x184200},
	{"SPSR_abt", REG_SPSR_abt, 0x1c4320},
	{"SPSR_EL1", REG_SPSR_EL1, 0x184000},
	{"SPSR_fiq", REG_SPSR_fiq, 0x1c4360},
	{"SPSR_irq", REG_SPSR_irq, 0x1c4300},
	{"SPSR_und", REG_SPSR_und, 0x1c4340},
	{"SSBS", REG_SSBS, 0x1b42c0},
	{"TCO", REG_TCO, 0x1b42e0},
	{"TCR_EL1", REG_TCR_EL1, 0x182040},
	{"TFSR_EL1", REG_TFSR_EL1, 0x185600},
	{"TFSRE0_EL1", REG_TFSRE0_EL1, 0x185620},
	{"TPIDR_EL0", REG_TPIDR_EL0, 0x1bd040},
	{"TPIDR_EL1", REG_TPIDR_EL1, 0x18d080},
	{"TPIDRRO_EL0", REG_TPIDRRO_EL0, 0x1bd060},
	{"TRFCR_EL1", REG_TRFCR_EL1, 0x181220},
	{"TTBR0_EL1", REG_TTBR0_EL1, 0x182000},
	{"TTBR1_EL1", REG_TTBR1_EL1, 0x182020},
	{"UAO", REG_UAO, 0x184280},
	{"VBAR_EL1", REG_VBAR_EL1, 0x18c000},
	{"ZCR_EL1", REG_ZCR_EL1, 0x181200},
}

func SysRegEnc(r int16) (string, uint32) {
	// The automatic generator guarantees that the order
	// of Reg in SystemReg struct is consistent with the
	// order of system register declarations
	if r <= AUTO_SYSREG_BEGIN || r >= AUTO_SYSREG_END {
		return "", 0
	}
	v := SystemReg[r-AUTO_SYSREG_BEGIN-1]
	return v.Name, v.Enc
}
