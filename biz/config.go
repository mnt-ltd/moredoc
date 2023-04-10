package biz

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"
	"moredoc/util/device"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConfigAPIService struct {
	pb.UnimplementedConfigAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewConfigAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *ConfigAPIService) {
	return &ConfigAPIService{dbModel: dbModel, logger: logger.Named("ConfigAPIService")}
}

func (s *ConfigAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// UpdateConfig 更新配置
func (s *ConfigAPIService) UpdateConfig(ctx context.Context, req *pb.Configs) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	var cfgs []*model.Config
	err = util.CopyStruct(req.Config, &cfgs)
	if err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("req", req), zap.Any("cfgs", cfgs), zap.Error(err))
		fmt.Println(err.Error())
	}

	doesUpdateSEO := false
	isEmail := false
	for idx, cfg := range cfgs {
		if cfg.Category == model.ConfigCategoryEmail && cfg.Name == model.ConfigEmailPassword && cfg.Value == "******" {
			// 6个星号，不修改密码
			cfgs[idx].Value = s.dbModel.GetConfigOfEmail(model.ConfigEmailPassword).Password
		}
		isEmail = isEmail || cfg.Category == model.ConfigCategoryEmail
		if cfg.Category == model.ConfigCategorySystem {
			doesUpdateSEO = true
		}
	}

	err = s.dbModel.UpdateConfigs(cfgs, "value")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if isEmail {
		cfgEmail := s.dbModel.GetConfigOfEmail(model.ConfigEmailEnable, model.ConfigEmailTestEmail)
		if cfgEmail.Enable && cfgEmail.TestEmail != "" {
			err = s.dbModel.SendMail("测试邮件", cfgEmail.TestEmail, "这是一封测试邮件")
			if err != nil {
				return nil, status.Error(codes.Internal, "邮件发送失败:"+err.Error())
			}
		}
	}

	if doesUpdateSEO {
		s.dbModel.InitSEO()
	}

	return &emptypb.Empty{}, nil
}

// ListConfig 查询配置
func (s *ConfigAPIService) ListConfig(ctx context.Context, req *pb.ListConfigRequest) (*pb.Configs, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetConfigList{
		QueryIn: map[string][]interface{}{
			"category": util.Slice2Interface(req.Category),
		},
	}

	configs, err := s.dbModel.GetConfigList(opt)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbConfigs []*pb.Config
	util.CopyStruct(&configs, &pbConfigs)

	for idx, cfg := range pbConfigs {
		if cfg.Category == model.ConfigCategoryEmail && cfg.Name == model.ConfigEmailPassword {
			pbConfigs[idx].Value = "******"
		}
	}

	return &pb.Configs{Config: pbConfigs}, nil
}

// GetSettings 获取公开配置
func (s *ConfigAPIService) GetSettings(ctx context.Context, req *emptypb.Empty) (*pb.Settings, error) {
	res := &pb.Settings{
		// Captcha:  &pb.ConfigCaptcha{},
		System:   &pb.ConfigSystem{},
		Footer:   &pb.ConfigFooter{},
		Security: &pb.ConfigSecurity{},
		Display:  &pb.ConfigDisplay{},
	}

	// captcha := s.dbModel.GetConfigOfCaptcha()
	// util.CopyStruct(&captcha, res.Captcha)

	system := s.dbModel.GetConfigOfSystem()
	if err := util.CopyStruct(&system, res.System); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("system", system), zap.Any("res.System", res.System), zap.Error(err))
	}
	system.Analytics = strings.TrimSpace(system.Analytics)
	if system.Analytics != "" {
		gq, errGQ := goquery.NewDocumentFromReader(strings.NewReader(system.Analytics))
		if errGQ == nil {
			var texts []string
			gq.Find("script").Each(func(i int, selection *goquery.Selection) {
				if text := strings.TrimSpace(selection.Text()); text != "" {
					texts = append(texts, text)
				}
			})
			if len(texts) > 0 {
				res.System.Analytics = strings.Join(texts, "\n")
			}
		}
	}
	res.System.Version = util.Version
	res.System.CreditName = s.dbModel.GetConfigOfScore(model.ConfigScoreCreditName).CreditName
	footer := s.dbModel.GetConfigOfFooter()
	if err := util.CopyStruct(&footer, res.Footer); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("footer", footer), zap.Any("res.Footer", res.Footer), zap.Error(err))
	}

	security := s.dbModel.GetConfigOfSecurity()
	if err := util.CopyStruct(&security, res.Security); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("security", security), zap.Any("res.Security", res.Security), zap.Error(err))
	}

	display := s.dbModel.GetConfigOfDisplay()
	if err := util.CopyStruct(&display, res.Display); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("display", display), zap.Any("res.Display", res.Display), zap.Error(err))
	}

	return res, nil
}

func (s *ConfigAPIService) GetStats(ctx context.Context, req *emptypb.Empty) (res *pb.Stats, err error) {
	res = &pb.Stats{
		UserCount:       0,
		DocumentCount:   0,
		CategoryCount:   0,
		ArticleCount:    0,
		CommentCount:    0,
		BannerCount:     0,
		FriendlinkCount: 0,
		Os:              runtime.GOOS,
		Version:         util.Version,
		Hash:            util.Hash,
		BuildAt:         util.BuildAt,
	}
	res.Os, _ = util.GetOSRelease()
	res.UserCount, _ = s.dbModel.CountUser()
	res.UserCount += s.dbModel.GetConfigOfDisplay(model.ConfigDisplayVirtualRegisterCount).VirtualRegisterCount
	res.DocumentCount, _ = s.dbModel.CountDocument()
	_, errPermission := s.checkPermission(ctx)
	if errPermission == nil {
		res.CategoryCount, _ = s.dbModel.CountCategory()
		res.ArticleCount, _ = s.dbModel.CountArticle()
		res.CommentCount, _ = s.dbModel.CountComment()
		res.BannerCount, _ = s.dbModel.CountBanner()
		res.FriendlinkCount, _ = s.dbModel.CountFriendlink()
		res.ReportCount, _ = s.dbModel.CountReport()
	}
	return
}

// UpdateSitemap 更新站点地图
func (s *ConfigAPIService) UpdateSitemap(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.UpdateSitemap()
	if err != nil {
		s.logger.Error("UpdateSitemap", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *ConfigAPIService) GetEnvs(ctx context.Context, req *emptypb.Empty) (res *pb.Envs, err error) {
	res = &pb.Envs{}
	_, errPermission := s.checkPermission(ctx)
	if errPermission != nil {
		return
	}
	envs := []*pb.EnvDependent{
		{
			Name:        "LibreOffice",
			Description: "LibreOffice是由文档基金会开发的自由及开放源代码的办公套件。魔豆文库用于将office等文档转为pdf。",
			Cmd:         "soffice",
			IsRequired:  true,
		},
		{
			Name:        "Calibre",
			Description: "calibre是一个自由开源的电子书软件套装。魔豆文库用于将epub、mobi等电子书转为pdf。",
			Cmd:         "ebook-convert",
			IsRequired:  true,
		},
		{
			// mupdf
			Name:        "MuPDF",
			Description: "MuPDF是一款以C语言编写的自由及开放源代码软件库，是PDF和XPS解析和渲染引擎。魔豆文库用于将PDF转为svg、png等图片。",
			Cmd:         "mutool",
			IsRequired:  true,
		},
		{
			Name:        "SVGO",
			Description: "SVGO 是一个基于 Node.js 的工具，用于优化 SVG 矢量图形文件。魔豆文库用于压缩svg图片大小。",
			Cmd:         "svgo",
			IsRequired:  false,
		},
		{
			Name:        "PM2",
			Description: "PM2是JavaScript运行时Node.js的进程管理器。用于做魔豆文库的系统守护进程。Windows下建议使用PM2。",
			Cmd:         "pm2",
			IsRequired:  false,
		}, {
			Name:        "Supervisor",
			Description: "Supervisor是一个客户端/服务器系统，用于监视进程状态，当进程不再运行时自动重启它们。用于做魔豆文库的系统守护进程。Linux下建议使用Supervisor。",
			Cmd:         "supervisorctl",
			IsRequired:  false,
		},
	}
	for i := 0; i < len(envs); i++ {
		now := time.Now()
		err := util.CheckCommandExists(envs[i].Cmd)
		envs[i].IsInstalled = err == nil
		envs[i].CheckedAt = &now
		if err != nil {
			envs[i].Error = err.Error()
		}
	}
	res.Envs = envs
	return
}

func (s *ConfigAPIService) GetDeviceInfo(ctx context.Context, req *emptypb.Empty) (res *pb.DeviceInfo, err error) {
	res = &pb.DeviceInfo{
		Cpu:    &pb.CPUInfo{},
		Memory: &pb.MemoryInfo{},
	}
	cpu := device.GetCPU()

	err = util.CopyStruct(&cpu, res.Cpu)
	if err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("cpu", cpu), zap.Any("res.Cpu", res.Cpu), zap.Error(err))
		return
	}
	res.Cpu.Cores = int32(runtime.NumCPU())

	mem := device.GetMemory()
	err = util.CopyStruct(&mem, res.Memory)
	if err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("mem", mem), zap.Any("res.Memory", res.Memory), zap.Error(err))
		return
	}

	res.Memory.Free = res.Memory.Total - res.Memory.Used

	disks := device.GetDisk()
	if len(disks) > 0 {
		for _, disk := range disks {
			pbDisk := &pb.DiskInfo{}
			err = util.CopyStruct(&disk, pbDisk)
			if err != nil {
				s.logger.Error("util.CopyStruct", zap.Any("disk", disk), zap.Any("res.Disk", res.Disk), zap.Error(err))
				return
			}
			res.Disk = append(res.Disk, pbDisk)
		}
	}
	return res, nil
}
