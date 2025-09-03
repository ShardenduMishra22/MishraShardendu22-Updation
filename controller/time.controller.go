package controller

import (
	"github.com/MishraShardendu22/models"
	"github.com/MishraShardendu22/util"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func ExperienceTimeline(c *fiber.Ctx) error {
	var exps []models.Experience
	var vexps []models.VolunteerExperience

	if err := mgm.Coll(&models.Experience{}).SimpleFind(&exps, bson.M{}); err != nil {
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Failed to fetch experiences", nil, "")
	}

	if len(exps) == 0 {
		return util.ResponseAPI(c, fiber.StatusOK, "No experiences found", nil, "")
	}

	if err := mgm.Coll(&models.VolunteerExperience{}).SimpleFind(&vexps, bson.M{}); err != nil {
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Failed to fetch volunteer experiences", nil, "")
	}

	if len(vexps) == 0 {
		return util.ResponseAPI(c, fiber.StatusOK, "No volunteer experiences found", nil, "")
	}

	exps = ReverseExperiences(exps)
	vexps = ReverseVolunteerExperiences(vexps)

	var expTimeline []models.ExperienceTimeLine
	var vexpTimeline []models.VolunteerExperienceTimeLine

	for _, exp := range exps {
		expTimeline = append(expTimeline, exp.ExperienceTimeline...)
	}

	for _, vexp := range vexps {
		vexpTimeline = append(vexpTimeline, vexp.VolunteerTimeLine...)
	}

	// Combine both timelines or return them separately based on your requirements
	response := map[string]interface{}{
		"experience_timeline":           expTimeline,
		"volunteer_experience_timeline": vexpTimeline,
	}

	return util.ResponseAPI(c, fiber.StatusOK, "Timeline fetched successfully", response, "")
}
